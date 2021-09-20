package main

import (
	"archive/zip"
	"bytes"
	"encoding/base64"
	"fmt"

	"syscall/js"
)

func converter(this js.Value, inputs []js.Value) interface{} {
	files := inputs[0]

	inBuf := make([]uint8, files.Get("byteLength").Int())

	js.CopyBytesToGo(inBuf, files)

	buf := new(bytes.Buffer)

	// Create a new zip archive.
	zipWriter := zip.NewWriter(buf)

	zipFile, err := zipWriter.Create("image.png")

	if err != nil {
		errStr := fmt.Sprintf("unable to write image. Error %s occurred\n", err)
		result := map[string]interface{}{
			"error": errStr,
		}
		return result
	}
	_, err = zipFile.Write(inBuf)

	if err != nil {
		errStr := fmt.Sprintf("unable to write zipfile. Error %s occurred\n", err)
		result := map[string]interface{}{
			"error": errStr,
		}
		return result
	}
	err = zipWriter.Close()

	if err != nil {
		errStr := fmt.Sprintf("unable to close zipWriter. Error %s occurred\n", err)
		result := map[string]interface{}{
			"error": errStr,
		}
		return result
	}

	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func main() {
	c := make(chan bool)
	js.Global().Set("convert", js.FuncOf(converter))
	<-c
}
