package main

import (
	"fmt"
	"math"
)

func main() {

	var (
		maxlength float64 = -1
		str       string  = "abacb"
	)

	if len(str) == 0 {
		maxlength = 0
	} else if len(str) == 1 {
		maxlength = 1
	}

	text := ""

	for _, val := range str {

		con, ind := find(text, val)

		if con {
			text = text[ind+1:]
		}

		text += string(val)
		maxlength = math.Max(float64(len(text)), maxlength)
	}

	fmt.Println(maxlength)
}

func find(str string, ch rune) (bool, int) {
	for index, val := range str {
		if val == ch {
			return true, index
		}
	}
	return false, 0
}
