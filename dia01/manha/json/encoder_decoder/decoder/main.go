package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"strings"
)

/* ETL -> Extract/Transform/Load  */
/* read -> transform -> write  */

func main() {

	const jsonStream = `
    	{"ProductID": "AXW123", "Price": 25.50}
    	{"ProductID": "NLBR17", "Price": 357.58}
    	{"ProductID": "KNUB82", "Price": 150}
    `

	myStreaming := strings.NewReader(jsonStream)
	myDecoder := json.NewDecoder(myStreaming)
	type MyData struct {
		ProductID string
		Price     float64
	}
	for {
		var data MyData
		if err := myDecoder.Decode(&data); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Data:", data.ProductID, data.Price)
	}

}

// versionamento semantico
// major.minor.patch
// 2.1.0
