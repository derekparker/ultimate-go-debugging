package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/go-delve/delve/service/rpc2"
)

func main() {
	if len(os.Args) < 2 {
		bail("Not enough arguments passed in, please provide address to connect to.")
	}

	serverAddr := os.Args[1]

	// Create a new connection to the Delve debug server.
	// rpc2.NewClient will log.Fatal if connection fails so there
	// won't be an error to handle here.
	client := rpc2.NewClient(serverAddr)

	defer client.Disconnect(true)

	// Stop the program we are debugging.
	// The act of halting the program will return it's current state.
	state, err := client.Halt()
	if err != nil {
		bail(err)
	}

	// Continue the program.
	client.Continue()

	// Write program status to stdout.
	enc := json.NewEncoder(os.Stdout)
	enc.Encode(state)
}

func bail(s interface{}) {
	fmt.Println(s)
	os.Exit(1)
}