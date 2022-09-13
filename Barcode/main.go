package main

import (
	"image/png"
	"log"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

func main() {

	// data that will be encoded into our barcode
	toEncode := "I love Go"
	// Generate a new writer for Code 128 barcode
	// this format allows you to encode all ASCII characters!
	writer := oned.NewCode128Writer()
	// with the writer, we can start encoding!
	img, err := writer.Encode(toEncode, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)
	if err != nil {
		log.Fatalf("impossible to encode barcode: %s", err)
	}
	// create a file that will hold our barcode
	file, err := os.Create("barcode.png")
	if err != nil {
		log.Fatalf("impossible to create file: %s", err)
	}
	defer file.Close()
	// Encode the image in PNG
	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("impossible to encode barcode in PNG: %s", err)
	}
}
