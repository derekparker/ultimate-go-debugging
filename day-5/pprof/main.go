package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

func main() {
	go func() { log.Println(http.ListenAndServe("localhost:6060", nil)) }()
	var wg sync.WaitGroup
	wg.Add(1)
	go leakyFunction(&wg)
	wg.Wait()
}

func leakyFunction(wg *sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 10000000; i++ {
		s = append(s, "debug this!!")
		if (i % 100000) == 0 {
			time.Sleep(500 * time.Millisecond)
		}
	}
}
