package main

import "fmt"

func main() {
	arr := []int{10, 15, 3, 7}

	val := 17

	dict := make(map[int]int)

	for _, k := range arr {

		if dict[k] == 0 {

			dict[k] = val - k

		}

		if k == dict[val-k] {

			fmt.Println(k, dict[k])
		}
	}
}
