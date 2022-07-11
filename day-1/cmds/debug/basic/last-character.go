package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello Gophers! Pass me a word and I will print the last character!")

	arg := os.Args[1]

	fmt.Printf("The last character in that word is: %c\n", arg[len(arg)])
}