package main

import (
	"fmt"
)

var (
	arr             = []int32{1, 3, 5, 4, 2}
	k, i, res int32 = 2, 0, 0
	arrLen          = int32(len(arr))
	dict            = make(map[int32]int32)
)

func main() {

	for ; i < arrLen; i++ {

		if k < arr[i] {
			_, ok := dict[arr[i]]
			if !ok {
				dict[arr[i]] += 1
			} else {
				dict[arr[i]]++
				res++
			}
		} else {
			_, ok := dict[arr[i]-k]
			if !ok {
				dict[arr[i]-k] += 1
			} else {
				res += dict[arr[i]-k]
			}
		}
	}

	fmt.Println(res)
}
