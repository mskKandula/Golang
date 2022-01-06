package main

import (
	"fmt"
)

type Blocks struct {
	gym    bool
	school bool
	store  bool
}

func main() {
	dict := map[string]int{"gym": -1, "school": -1, "store": -1}
	blocks := []Blocks{
		{gym: false, school: true, store: false},
		{gym: true, school: false, store: false},
		{gym: true, school: true, store: false},
		{gym: false, school: true, store: false},
		{gym: false, school: true, store: true},
	}

	for index, block := range blocks {
		if block.gym && block.school && block.store {
			fmt.Println(index)
			break
		}
		if block.gym {
			dict["gym"] = index
		}

		if block.school {
			dict["school"] = index
		}

		if block.store {
			dict["store"] = index
		}

		if val := dict["gym"]; val >= 0 {
			if val := dict["school"]; val >= 0 {
				if val := dict["store"]; val >= 0 {
					min, max := getMinMax([]int{dict["school"], dict["gym"], dict["store"]})
					fmt.Println((min + max) / 2)
					break
				}
			}
		}

	}
}

func getMinMax(arr []int) (int, int) {
	min := arr[0]
	max := arr[0]
	for _, v := range arr {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	return min, max
}
