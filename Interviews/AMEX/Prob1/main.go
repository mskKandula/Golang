package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	number := 1254 // Change this to any number you'd like

	// Check if the number is a single digit
	if number < 10 {
		fmt.Println("Number is a single digit, returning 0")
		return
	}

	digits := len(strconv.Itoa(number))

	// Calculate the smallest number with the same number of digits
	smallestNumber := int(math.Pow10(digits - 1))

	fmt.Println("Smallest number with the same number of digits:", smallestNumber)
}
