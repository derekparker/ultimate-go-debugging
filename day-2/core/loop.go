package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for {
		fmt.Println("Loop iteration:", i)
		i++
		time.Sleep(500 * time.Millisecond)
	}
}
