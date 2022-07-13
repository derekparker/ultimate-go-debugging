package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	pid := os.Getpid()

	buf := make([]byte, 64)
	f, err := os.Open("/dev/random")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	_, err = f.Read(buf)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	t := time.Now()

	fmt.Println(pid, t)
}