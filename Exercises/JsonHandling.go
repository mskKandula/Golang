package main

import (
	"encoding/json"
	"fmt"
)

type Action struct {
	Type    string
	Payload string
}

type Lake struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Area int32  `json:"area"`
}

func main() {
	jsonHandler()
	jsonHandlerVariation()
}

func jsonHandler() {
	req := `{
		"type":"Post",
		"payload":"1"
	}`

	var action Action

	json.Unmarshal([]byte(req), &action)

	fmt.Println(action.Payload)

}

func jsonHandlerVariation() {

	req := `{
		"type":"Post",
		"payload":{
			"id":"1",
			"name":"Great Lake",
			"area":3200
		}
	}`
	lake := struct {
		Type    string `json:"type"`
		Payload Lake   `json:"payload"`
	}{}

	json.Unmarshal([]byte(req), &lake)

	fmt.Println(lake)

}
