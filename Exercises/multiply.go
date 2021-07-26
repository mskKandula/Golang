package main

import "fmt"

func mul(arr []int) int {

	prod := 1
	for _, k := range arr {
		prod = prod * k
	}

	return prod
}

func main() {

	arr := [5]int{1, 2, 3, 4, 5}

	krr := [5]int{}

	length := len(arr)

	if length == 0 || length == 1 {
		fmt.Println(arr)
	}

	for i := range arr {

		if i == 0 {

			krr[i] = mul(arr[1:length])

		} else if 0 < i && i < length-1 {

			krr[i] = mul(arr[0:i]) * mul(arr[i+1:length])

		} else {

			krr[i] = mul(arr[0 : length-1])
		}

	}
	fmt.Println(krr)

}
