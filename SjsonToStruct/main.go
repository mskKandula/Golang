package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/tidwall/sjson"
)

type Details struct {
	Name string `json:"name"`
	Age  string `json:"age"`
}

func main() {
	str := `{}`

	str, _ = sjson.Set(str, "name", "abc")
	str, _ = sjson.Set(str, "age", "123")

	var details Details

	json.Unmarshal([]byte(str), &details)

	fmt.Println(reflect.TypeOf(details.Age))

}
