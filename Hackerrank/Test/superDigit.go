package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
	t1 := time.Now()
	fmt.Println(superDigit("98963", 100))
	fmt.Println("Time taken---", time.Since(t1))
}

func superDigit(n string, k int32) int32 {
	arr := strings.Split(n, "")
	init := ""
	num := 0
	for i := range arr {
		res, _ := strconv.Atoi(arr[i])
		num += res
	}
	init = fmt.Sprintf("%v", num*int(k))

	for {
		if len(init) == 1 {
			break
		}
		alpha := []rune(init)
		result := int32(0)
		for i := 0; i < len(init); i++ {
			if !unicode.IsDigit(alpha[i]) {
				return 0
			} else {
				res, _ := strconv.Atoi(fmt.Sprintf("%c", alpha[i]))
				result += int32(res)
			}
		}
		init = fmt.Sprintf("%v", result)
	}

	res, _ := strconv.Atoi(init)
	return int32(res)
}
