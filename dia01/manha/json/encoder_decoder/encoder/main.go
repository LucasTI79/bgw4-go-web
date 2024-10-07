package main

import (
	"encoding/json"
	"os"
)

func main() {
	myEncoder := json.NewEncoder(os.Stdout)

	// prepare the information you want to send in json format to the streaming
	type MyData struct {
		ProductID string
		Price     float64
	}

	data := MyData{
		ProductID: "XASD",
		Price:     25.50,
	}
	// the Encode method is invoked.
	// internally this method makes a kind of marshall and writes it to the stream

	myEncoder.Encode(data)
}
