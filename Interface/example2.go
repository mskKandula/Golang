package main

import (
	"fmt"
)

type a struct {
	n int
}

type b struct {
	s string
}

type printOut interface {
	call()
}

func (example a) call() {
	fmt.Println(example.n)

}

func (example b) call() {
	fmt.Println(example.s)

}

func print(p printOut) {
	p.call()
}

func main() {

	// var name b

	// name.s = "abc"

	// print(name)

	var number a

	number.n = 123

	print(number)
}
