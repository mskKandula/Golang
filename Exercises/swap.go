package main

import (
	"fmt"
	"sort"
)

func split(num int) []int {
	var arr []int
	for num > 0 {
		last := num % 10
		arr = append(arr, last)
		num = num / 10
	}
	sort.Ints(arr)
	return arr

}
func main() {
	num := 5564
	arr := split(num)
	fmt.Println(arr)
}
