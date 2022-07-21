package main

/*
int callme(int i) {
	return 42 + i;
}
*/
import "C"

func main() {
	println(C.callme(42))
}
