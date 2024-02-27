package main

import (
	"encoding/json"
	"fmt"
)

type Product struct {
	ProductId   int    `json:"productId"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func main() {

	var products []Product

	discountDict := make(map[int]int)

	discountDict[1] = 10
	discountDict[2] = 20
	discountDict[3] = 30

	str := `[{
        "productId":1,
        "name": "golang",
        "price":18,
        "description":"Programming Language"

    },
    {
        "productId":2,
        "name": "python",
        "price":15,
        "description":"Programming Language"
    },{
        "productId":3,
        "name": "Javascript",
        "price":12,
        "description":"Programming Language"
    }]`

	err := json.Unmarshal([]byte(str), &products)
	if err != nil {
		fmt.Println(err)
	}

	for i := range products {
		products[i].ApplyDiscount(discountDict)
	}

	fmt.Println(products)

}

func (product *Product) ApplyDiscount(discountDict map[int]int) {

	discount := discountDict[product.ProductId]

	discountPrice := (product.Price * discount / 100)

	product.Price = product.Price - discountPrice

}

// set of products [{},{}]
