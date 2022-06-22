package main

import (
	"fmt"
	"reflect"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func main() {
	str := `{}`

	str, _ = sjson.Set(str, "name", "abc")
	str, _ = sjson.Set(str, "age", "123")
	fmt.Println(str)
	fmt.Println(reflect.TypeOf(str))

	strParse := gjson.Parse(str)
	fmt.Println(strParse)
	fmt.Println(reflect.TypeOf(&strParse))
	fmt.Println(`{"golang":{` + strParse.String() + `}`)
	fmt.Println(reflect.TypeOf(strParse.String()))
	fmt.Println(reflect.TypeOf(strParse.Value()))
}
