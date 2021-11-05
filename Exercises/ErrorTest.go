package main

import (
	"encoding/json"
	"fmt"
	"log"
)

var obj map[string]interface{}

func main() {
	jsonHandler()
	jsonHandlerVariation()
}

func jsonHandler() {
	req := `{
		"type":"Post",
		"payload":"1"
	}`

	if err := json.Unmarshal([]byte(req), &obj); err != nil {
		log.Fatal(err)
	}

	fmt.Println(obj["payload"])

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

	if err := json.Unmarshal([]byte(req), &obj); err != nil {
		log.Fatal(err)
	}

	fmt.Println(obj["payload"])
}
