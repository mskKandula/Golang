package main

import (
	"fmt"
	"strconv"
)

func main() {

	binaryNums := []string{"0", "100", "101", "01", "010"}
	str, err := returnMissingNum(binaryNums)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(str)
}

func returnMissingNum(binNums []string) (string, error) {
	var (
		found int64
		res   string
		nums  [6]int
	)

	for _, val := range binNums {

		num, err := strconv.ParseInt(val, 2, 64)
		if err != nil {
			return "", err
		}

		nums[num] = 1
	}

	for i, val := range nums {
		if val != 1 {
			found = int64(i)
		}
	}

	res = strconv.FormatInt(found, 2)
	return res, nil
}
