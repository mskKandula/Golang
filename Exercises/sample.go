package main

import (
	"fmt"

	"github.com/sindbach/json-to-bson-go/convert"
	"github.com/sindbach/json-to-bson-go/options"
)

func main() {
	doc := `{"foo": "buildfest", "bar": {"$numberDecimal":"2021"} }`
	opt := options.NewOptions()
	result, _ := convert.Convert([]byte(doc), opt)
	fmt.Println(result)
}
