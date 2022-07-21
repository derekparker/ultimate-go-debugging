package main

/*
#include "cstuff.h"
*/
import "C"

func main() {
	println(C.callme(42))
	println(callmefromgo(42))
}

func callmefromgo(i int) int {
	return 42 + i
}
