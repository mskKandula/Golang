package main

import (
	"encoding/json"
	"fmt"
	"runtime"
	"strings"
	"time"
)

const testBytes = `{ "Test": "value" }`

type Message struct {
	Test string
}

func cpuIntensive(p *Message) {
	for i := int64(1); i <= 1000; i++ {
		json.NewDecoder(strings.NewReader(testBytes)).Decode(p)
	}
	fmt.Println("Done intensive thing")
}

func printVar(p *Message) {
	fmt.Printf("print x = %v.\n", *p)
}

func main() {
	runtime.GOMAXPROCS(2)

	x := Message{}
	go cpuIntensive(&x) // This should go into background
	go printVar(&x)

	// This won't get scheduled until everything has finished.
	time.Sleep(5 * time.Second) // Wait for goroutines to finish
}
