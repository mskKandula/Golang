package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Inner struct {
	Barcode   string  `json:"barcode"`
	Item      string  `json:"item"`
	Category  string  `json:"category"`
	Price     float64 `json:"price"`
	Discount  float64 `json:"discount"`
	Available int     `json:"available"`
}

type Result struct {
	Page        int     `json:"page"`
	Per_page    int     `json:"per_page"`
	Total       int     `json:"total"`
	Total_pages int     `json:"total_pages"`
	Data        []Inner `json:"data"`
}

func main() {

	fmt.Println(getProductsInRange("Accessories", 100, 3000))

}

func getProductsInRange(category string, minPrice int32, maxPrice int32) int32 {

	var (
		result Result
		count  int32
		page   int = 1
	)

	for {
		url := fmt.Sprintf("https://jsonmock.hackerrank.com/api/inventory?category=%s&page=%d", category, page)

		resp, err := http.Get(url)
		if err != nil {
			log.Panic(err)
		}

		defer resp.Body.Close()

		json.NewDecoder(resp.Body).Decode(&result)
		// fmt.Println(result)

		for _, res := range result.Data {
			if minPrice <= int32(res.Price) && int32(res.Price) <= maxPrice {
				count += 1
			}
		}

		if page == result.Total_pages {
			break
		}

		page += 1

	}

	return count
}
