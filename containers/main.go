package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server listening at :8080")
	http.HandleFunc("/", helloServer)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	if err != nil {
		panic(err)
	}
}