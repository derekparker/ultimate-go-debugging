package main

import (
	"fmt"
)

func main() {
	sayHello()
	gatherSentiment()
}

func sayHello() {
	str := "Hello, Gophers!"
	fmt.Println(str)
}

func gatherSentiment() {
	fmt.Println("How would you rate the class from 1-10:")

	var i int
	_, err := fmt.Scanln(&i)
	if err != nil {
		panic(err)
	}

	if !withinBounds(&i) {
		panic("user input not within expected bounds")
	}

	switch {
	case i > 0 && i < 3:
		fmt.Println("Guess it's going terribly... :(")
	case i > 2 && i < 7:
		fmt.Println("Guess it's going alright... :-|")
	case i > 6 && i < 11:
		fmt.Println("Guess it's going GREAT!!! :-D")
	default:
		panic(fmt.Sprintf("Invalid response received: %d\n", i))
	}
}

func withinBounds(n *int) bool {
	*n -= 10
	return *n <= 0
}