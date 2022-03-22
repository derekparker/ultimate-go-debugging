package main

import (
	"time"
)

func main() {
	var i *int
	for {
		*i++
		time.Sleep(500 * time.Millisecond)
	}
}