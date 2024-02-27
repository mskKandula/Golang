package main

import "fmt"

func main() {
	fmt.Println(Solution("aabb"))
}

func Solution(S string) string {
	var occurrences [26]int
	for _, ch := range S {
		occurrences[int(ch)-int('a')]++
	}

	fmt.Println(occurrences)

	var best_char rune = 'a'
	var best_res int = 0

	for i := 0; i < 26; i++ {
		if occurrences[i] > best_res {
			best_char = rune(int('a') + i)
			best_res = occurrences[i]
		}
	}

	return string(best_char)
}
