package main

import (
	"encoding/json"
	"fmt"
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

func main() {
	myRecord := Record{
		Name:    "Ivan",
		Surname: "Petrov",
		Tel: []Telephone{
			{Mobile: true, Number: "+79991234455"},
			{Mobile: true, Number: "+79991231155"},
			{Mobile: false, Number: "+79991564155"},
		},
	}

	rec, err := json.Marshal(&myRecord)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(rec))

	var unRec Record
	err1 := json.Unmarshal(rec, &unRec)
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	fmt.Println(unRec)

}
