package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Test struct {
	Name    string   `json:"name"`
	Weather string   `json:"weather"`
	Status  []string `json:"status"`
}

type DataCode struct {
	Data []Test `json:"data"`
}

func main() {

	getTemperature("dallas")

}

func getTemperature(name string) {

	var dataCode DataCode

	url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/weather?name=%s", name)

	resp, _ := http.Get(url)

	json.NewDecoder(resp.Body).Decode(&dataCode)

	fmt.Println(dataCode.Data[0])

}
