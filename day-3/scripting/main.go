package main

import (
	"bytes"
	"context"
	"encoding/binary"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

// packet is a small struct that contains an id and associated value.
type packet struct {
	id  int
	val []byte
}

func produceValues(ctx context.Context, wg *sync.WaitGroup, ch chan<- packet) {
	// Let WaitGroup know we're finished executing.
	defer wg.Done()
	// Create a bytes buffer of length 64.
	buf := make([]byte, 1<<6)

	for {
		// Create a random ID.
		id := rand.Int()
		// Read some random data into our buffer.
		if _, err := rand.Read(buf); err != nil {
			panic(err)
		}
		// Create a packet struct to send over the channel.
		p := packet{id: id, val: buf}
		select {
		case <-ctx.Done():
			// If the context has been cancelled, return.
			return
		case ch <- p:
			// Do nothing.
		}

		// Sleep 0.5 seconds.
		time.Sleep(10 * time.Millisecond)
	}
}

func consumeValues(ctx context.Context, wg *sync.WaitGroup, ch <-chan packet) {
	// Let WaitGroup know we're finished executing.
	defer wg.Done()

	// Create buffer to store values in-memory.
	buf := make([]packet, 0)
	for {
		select {
		case <-ctx.Done():
			// If the context has been cancelled, return.
			return
		case pkt := <-ch:
			// Store data in our buffer.
			buf = append(buf, pkt)

			fmt.Println("got value:", binary.LittleEndian.Uint64(pkt.val))

			if bytes.Contains(pkt.val, []byte{0xcc, 0xff}) {
				fmt.Printf("got invalid packet with ID: %d\n", pkt.id)
			}
		}
	}
}

func main() {
	// Create a WaitGroup to ensure program does not exit before goroutines finish.
	var wg sync.WaitGroup
	// Create a context object to pass to our goroutines.
	ctx, cancel := context.WithCancel(context.Background())

	// Create a channel to be notified when we get a signal.
	sigch := make(chan os.Signal, 1)
	// Ensure we are notified on receipt of SIGTERM.
	signal.Notify(sigch, syscall.SIGTERM)

	go func(sch <-chan os.Signal, cncl func()) {
		// Wait until we get a signal.
		<-sch

		fmt.Println("Got interrupt signal, exiting...")

		// Call the context cancellation function.
		cncl()
	}(sigch, cancel) // Pass along arguments to prevent closing over variables.

	// Create a channel to pass along values between goroutines.
	ch := make(chan packet)

	// Add 2 to our WaitGroup since we are starting 2 goroutines.
	wg.Add(2)

	// Begin executing goroutine to produce values.
	go produceValues(ctx, &wg, ch)

	// Begin executing goroutine to consume values.
	go consumeValues(ctx, &wg, ch)

	// Wait until our goroutines have finished.
	wg.Wait()
}