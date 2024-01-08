package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	// Create the target file
	out, err := os.Create("target.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	// Open the video file
	video, err := os.Open("golang.mp4")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer video.Close()

	// Open the pdf file
	pdf, err := os.Open("msk.pdf")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer pdf.Close()

	// Copy the video file to the target file
	if _, err := io.Copy(out, video); err != nil {
		fmt.Println(err)
		return
	}

	// Copy the pdf file to the target file
	if _, err := io.Copy(out, pdf); err != nil {
		fmt.Println(err)
		return
	}

	// Create a buffer for the video file
	buf := bytes.NewBuffer(nil)

	// Copy the video file to the buffer
	if _, err := io.Copy(buf, video); err != nil {
		fmt.Println(err)
		return
	}

	// Write the buffer to the target file
	if _, err := io.Copy(out, buf); err != nil {
		fmt.Println(err)
		return
	}
}
