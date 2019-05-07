package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Name string
	Body string
	Time int64
}

func main() {
	m := Message{"Alice", "Hello", 1294706395881547000}

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(b)

	var m1 Message
	err1 := json.Unmarshal(b, &m1)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(m1)

	b1 := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	var m2 Message
	err2 := json.Unmarshal(b1, &m2)
	if err2 != nil {
		fmt.Println(err2)
	}
	fmt.Println(m2)
}
