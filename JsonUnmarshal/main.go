package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cakeType struct {
	ID      string  `json:"id"`
	Type    string  `json:"type"`
	Name    string  `json:"name"`
	Ppu     float64 `json:"ppu"`
	batters `json:"batters"`
	Topping []topping `json:"topping"`
}
type batter struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
type batters struct {
	Batter []batter `json:"batter"`
}
type topping struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

func main() {
	str := `{
		"id": "0001",
		"type": "Cake",
		"ppu": 0.55,
		"batters":
			{
				"batter":
					[
						{ "id": "1001", "type": "Regular" },
						{ "id": "1002", "type": "Chocolate" },
						{ "id": "1003", "type": "Blueberry" },
						{ "id": "1004", "type": "Devil's Food" }
					]
			},
		"topping":
			[
				{ "id": "5001", "type": "None" },
				{ "id": "5002", "type": "Glazed" },
				{ "id": "5005", "type": "Sugar" },
				{ "id": "5007", "type": "Powdered Sugar" },
				{ "id": "5006", "type": "Chocolate with Sprinkles" },
				{ "id": "5003", "type": "Chocolate" },
				{ "id": "5004", "type": "Maple" }
			]
	}`

	var cake cakeType

	err := json.Unmarshal([]byte(str), &cake)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cake)

}
