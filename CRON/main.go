package main

import (
	"fmt"
	"time"
)

func call(f func(), n int) {
	for {
		f()
		time.Sleep(time.Duration(n) * time.Second)
	}
}

func printHello() {
	fmt.Println("hello")
}

func main() {
	call(printHello, 5)
}
