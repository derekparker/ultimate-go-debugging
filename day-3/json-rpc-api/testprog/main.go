package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	setupHandlers()
	startServer()
}

func setupHandlers() {
	http.HandleFunc("/", echoHandler)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
	if err != nil {
		panic(err)
	}
}

func startServer() {
	fmt.Println("Server started on :8081")
	log.Fatal(http.ListenAndServe(":8082", nil))
}
