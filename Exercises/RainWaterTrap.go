package main

import "fmt"

func main() {
	var (
		arr            = []int{3, 0, 0, 2, 0, 4}
		max_val, count int
	)

	for index, val := range arr {
		if index == 0 {
			max_val = val
			continue
		}

		if max_val-val > 0 {
			count += max_val - val
		} else {
			max_val = val
		}

	}

	fmt.Println(count)

}
