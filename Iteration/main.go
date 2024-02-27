package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	fakeData := `[{"name":"abc","age":10},
{"name":"abc","age":20},
{"name":"abc","age":30},
{"name":"abc","age":40},
{"name":"abc","age":50},
{"name":"abc","age":60},
{"name":"abc","age":70}]`

	var persons []Person
	json.Unmarshal([]byte(fakeData), &persons)
	for _, person := range persons {
		fmt.Println(person.Age)
	}
}
