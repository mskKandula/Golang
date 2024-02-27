package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World")
	data := callToDB()
	file := prepareFile(data)
	uploadFile(file)
}

func callToDB() {
	time.Sleep(100 * time.Millisecond)

}

func prepareFile() {

}

func uploadFile() {

}
