package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Record struct {
	Name    string
	Surname string
	Tel     []Telephone
}

type Telephone struct {
	Mobile bool
	Number string
}

func saveToJSON(filename *os.File, key interface{}) {
	encodeJSON := json.NewEncoder(filename)
	err := encodeJSON.Encode(key)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func main() {
	myRecord := Record{
		Name:    "Ivan",
		Surname: "Ivanov",
		Tel: []Telephone{
			{Mobile: true, Number: "1234-1223"},
			{Mobile: true, Number: "1234-1243"},
			{Mobile: false, Number: "1234-1213"},
		},
	}

	saveToJSON(os.Stdout, myRecord)
}
