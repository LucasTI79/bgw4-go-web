package main

import (
	"fmt"
	"os"
)

func main() {
	filePath := "./.env"
	// open file
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	port := os.Getenv("PORT")
	fmt.Println(port)
}
