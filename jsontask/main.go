package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName  string `json:",omitempty"`
	LastName   string `json:",omitempty"`
	Age        int    `json:",omitempty"`
	Address    string `json:",omitempty"`
	IsMarriege bool   `json:",omitempty"`
}

func main() {
	p := Person{"Ivan", "Petrov", 25, "Sofia, Triadista 155", false}

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	var p1 Person
	err1 := json.Unmarshal(b, &p1)
	if err != nil {
		fmt.Println(err1)
	}
	fmt.Println(p1)

	p2 := Person{FirstName: "Ivan", Age: 25}
	b1, err2 := json.Marshal(p2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(string(b1))
}
