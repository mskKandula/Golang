package main

import (
	"encoding/json"
	_ "image/jpeg"
	_ "image/png"
	"syscall/js"

	"github.com/subeshb1/wasm-go-image-to-ascii/convert"
)

func converter(this js.Value, inputs []js.Value) interface{} {
	imageArr := inputs[0]
	options := inputs[1].String()
	inBuf := make([]uint8, imageArr.Get("byteLength").Int())
	js.CopyBytesToGo(inBuf, imageArr)
	convertOptions := convert.Options{}
	err := json.Unmarshal([]byte(options), &convertOptions)
	if err != nil {
		convertOptions = convert.DefaultOptions
	}
	converter := convert.NewImageConverter()
	return converter.ImageFile2ASCIIString(inBuf, &convertOptions)
}
func main() {
	c := make(chan bool)
	js.Global().Set("convert", js.FuncOf(converter))
	<-c
}
