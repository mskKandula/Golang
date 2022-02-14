package main

import "fmt"

func main() {
	var (
		arr []int64
		r   int64
	)
	arr, r = []int64{1, 4, 16, 64}, 4
	res := countTriplets(arr, r)
	fmt.Println(res)
}

func countTriplets(arr []int64, r int64) int64 {
	v2 := make(map[int64]int64)
	v3 := make(map[int64]int64)
	var total int64

	for _, v := range arr {

		if val, ok := v3[v]; ok {
			total += val
			fmt.Println("24", val)
		}

		if val, ok := v2[v]; ok {
			v3[v*r] += val
			fmt.Println("29", val)
		}
		v2[v*r]++
		fmt.Println("32", v2[v*r])
	}
	return total

}
