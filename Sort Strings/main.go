package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello")

	fmt.Println(sortStrings([]string{"abc", "def"}))
}

func sortStrings(strs []string) []string {

	res := make([]string, len(strs), len(strs))
	for _, str := range strs {
		if len(str) <= 0 {
			continue
		}

		var result string

		for _, s := range str {
			result = string(s) + result
		}

		res = append(res, result)

	}
	return res
}
