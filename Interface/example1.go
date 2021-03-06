package main

import "fmt"

type number struct {
	n int
}

type name struct {
	s string
}

type printOut interface {
	print()
}

func (example name) print() {
	fmt.Println(example.s)
}

func (example number) print() {
	fmt.Println(example.n)
}

func main() {
	var p printOut

	p = name{"abc"}

	p.print()

}
