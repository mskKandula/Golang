package main

import (
	"fmt"
	"math/rand"
	"time"
)

type D struct {
	momentData    int
	doubleChannel chan chan int
}

func (d *D) getData() int {
	responseChan := make(chan int)
	d.doubleChannel <- responseChan
	response := <-responseChan
	return response
}

func (d *D) run() {

	seed := rand.NewSource(time.Now().UnixNano())
	gen := rand.New(seed)

	ticker := time.NewTicker(1 * time.Second)

	for {
		select {
		case <-ticker.C:
			d.momentData = gen.Int()

		case responseChan := <-d.doubleChannel:
			responseChan <- d.momentData

		}

	}

}

func main() {
	d := &D{
		doubleChannel: make(chan chan int),
	}

	go d.run()

	time.Sleep(1 * time.Second)

	fmt.Println(d.getData())
}
