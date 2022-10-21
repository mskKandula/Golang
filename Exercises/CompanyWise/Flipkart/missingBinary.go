package main

import (
	"fmt"
	"strconv"
)

func main() {

	binaryNums := []string{"0", "01", "010", "100", "101"}
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
	)

	for i, val := range binNums {

		num, err := strconv.ParseInt(val, 2, 64)
		if err != nil {
			return "", err
		}

		if i != int(num) {
			found = int64(i)
			break
		}
	}

	res = strconv.FormatInt(found, 2)
	return res, nil
}
