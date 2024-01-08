package main

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"os"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/segment"
	gosseract "github.com/otiai10/gosseract/v2"
)

func main() {
	fmt.Println("Hello")
	ocr("testocr.png")
}

func ocr(fileName string) {

	file, err := os.OpenFile(fileName, os.O_WRONLY, 0755)
	if err != nil {
		fmt.Println("20", err)
	}
	defer file.Close()

	// Decode file to image struct
	img, _, err := image.Decode(file)
	if err != nil {
		log.Fatalln(err)
	}

	// Convert Image to grayscale
	grayscale := effect.Grayscale(img)

	// Convert Image to threshold segment
	threshold := segment.Threshold(grayscale, 128)

	// Convert Image to Bytes
	buf := new(bytes.Buffer)
	jpeg.Encode(buf, threshold, nil)

	// Initiation Gosseract new client
	client := gosseract.NewClient()

	// close client when the main function is finished running
	defer client.Close()

	// Read byte to image and set whitelist character
	client.SetImageFromBytes(buf.Bytes())
	client.SetWhitelist(" -:/abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	// Get text result from OCR
	text, _ := client.Text()

	fmt.Println(text)
}
