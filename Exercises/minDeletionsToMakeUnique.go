package main

import (
	"fmt"
)

func main() {
	c := minDeletions("abbcccddd")
	fmt.Println(c)
}

func minDeletions(s string) int {
	dict := make(map[rune]int)
	dict2 := make(map[int]bool)
	count := 0
	if len(s) == 0 || len(s) == 1 {
		return 0
	} else {
		for _, letter := range s {
			dict[letter] += 1
		}

	}

	for key, val := range dict {
		if _, ok := dict2[val]; ok {
			ee := dict[key] - 1
			if _, ok = dict2[ee]; ok {
				count += val - 1
			} else {
				dict2[val-1] = true
				count += 1
			}

		} else {
			dict2[val] = true
		}
	}
	return count
}
