package main

import (
	"fmt"
)

func main() {

	//input: [1,2,3,4,5] sorted & unsorted, 7

	//output:4,3 or 5,2

	arr := []int{1, 2, 3, 4, 5}

	target := 7

	low, high := 0, len(arr)-1
	for low <= high {
		if arr[low]+arr[high] < target {
			low++
		} else if arr[low]+arr[high] > target {
			high--
		} else {
			fmt.Println(arr[low], arr[high])
			low++
			high--
		}
	}

	// var wg sync.WaitGroup

	// input := "Mohana"
	// ch := make(chan rune, 3)

	// wg.Add(2)

	// go func(str string) {
	// 	defer wg.Done()
	// 	for _, val := range str {
	// 		ch <- val
	// 	}

	// 	close(ch)

	// }(input)

	// go func() {
	// 	for data := range ch {
	// 		fmt.Println(string(data))
	// 	}
	// 	defer wg.Done()
	// }()

	// wg.Wait()
}

//2 goroutnes

//input string

//pass & print

//output

//user

//200
//200
//200
//500
//404
//

//post /api/o/v1/user/

//get /api/r/v1/user/:id

//get /api/r/v1/users

//put

//
