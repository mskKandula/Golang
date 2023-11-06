package main

import (
	"fmt"
	"reflect"

	"github.com/tidwall/gjson"
)

func main() {
	str := `["1","2","3"]`

	res := gjson.Parse(str).Array()

	for _, id := range res {
		fmt.Println("Hello World", reflect.TypeOf(id.String()))
	}

}
