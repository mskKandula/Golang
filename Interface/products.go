package main

import "fmt"

type Color int
type Size int

const (
	small Size = iota
	medium
	large
)

const (
	red Color = iota
	blue
	green
)

type Product struct {
	name  string
	color Color
	size  Size
}

type Filter struct {
}

func (f *Filter) FilterByColor(Products []Product, color Color) []*Product {
	result := make([]*Product, 0)

	for i, product := range Products {
		if product.color == color {
			result = append(result, &Products[i])

		}
	}
	return result
}

func (f *Filter) FilterBySize(Products []Product, size Size) []*Product {
	result := make([]*Product, 0)

	for i, product := range Products {
		if product.size == size {
			result = append(result, &Products[i])

		}
	}
	return result
}

func main() {
	var f Filter
	products := []Product{
		{"apple", green, small},
		{"apple", red, small},
		{"Watermelon", green, large},
	}

	// for _, product := range f.FilterByColor(products, green) {
	// 	fmt.Printf("-Green %s\n", product.name)
	// }

	for _, product := range f.FilterBySize(products, large) {
		fmt.Printf("-Large %s\n", product.name)
	}

}
