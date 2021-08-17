package main

import "fmt"

func main() {
	num := 3241
	count := checkKaprekars(num)
	fmt.Println(count)
}

func checkKaprekars(num int) int {
	count := 0

	if num <= 0 || num > 9999 {
		return count
	}

	for num != 6174 {
		minNum, maxNum := minMax(num)
		num = maxNum - minNum
		count++
	}
	return count
}

func minMax(num int) (int, int) {

	first := 0
	second := 0
	max := 0
	min := 0
	n := 0
	maxNum := 0
	minNum := 0

	for num != 0 {
		n = num % 10

		if n > max {
			min = second
			second = first
			first = max
			max = n
		} else if n > first {
			min = second
			second = first
			first = n
		} else if n > second {
			min = second
			second = n
		} else {
			min = n
		}
		num = num / 10
	}

	maxNum = 10*maxNum + max
	maxNum = 10*maxNum + first
	maxNum = 10*maxNum + second
	maxNum = 10*maxNum + min

	minNum = 10*minNum + min
	minNum = 10*minNum + second
	minNum = 10*minNum + first
	minNum = 10*minNum + max

	fmt.Printf("%v --- %v\n", maxNum, minNum)

	return minNum, maxNum
}
