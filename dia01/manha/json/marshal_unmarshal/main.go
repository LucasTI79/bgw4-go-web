package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	type product struct {
		Name      string `json:"name"`
		Price     int    `json:"price"`
		Published bool   `json:"published"`
	}

	p1 := product{
		Name:      "MacBook Pro",
		Price:     1500,
		Published: true,
	}

	jsonDataParsed, err := json.Marshal(p1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonDataParsed))

	jsonData := `{"Name": "MacBook Air", "Price": 900, "Published": true}`

	var p2 product

	if err := json.Unmarshal([]byte(jsonData), &p2); err != nil {
		log.Fatal(err)
	}

	fmt.Println(p2)
}
