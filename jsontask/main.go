package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName  string
	LastName   string
	Age        int
	Address    string
	IsMarriege bool
}

func main() {
	p := Person{"Ivan", "Petrov", 25, "Sofia, Triadista 155", false}

	b, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)


	var p1 Person
	err1:=json.Unmarshal(b,&p1)
	if err != nil{
		fmt.Println(err1)
	}
	fmt.Println(p1)
}
