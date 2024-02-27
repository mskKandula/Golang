// Problem Name: Palindrome
// Problem Statement: Given a string write a function returns true if that string is a palindrome
// Description: A Palindrome is a word that reads the same backward as forward
// Examples(s): "madam", "racecar".

// package main

// import "fmt"

// func main() {

// 	str := "madaam"
// 	start := 0
// 	end := len(str) - 1

// 	for start <= end {

// 		if str[start] == str[end] {
// 			start += 1
// 			end -= 1
// 			continue
// 		} else {
// 			fmt.Println("Not a palindrome")
// 			return
// 		}

// 	}

// 	fmt.Println("Palindrome")

// }

// Problem Name: Find Longest Palindromic Substring
// Problem Statement: Given a string s, find and return the longest palindromic substring in s. You may assume that the maximum length of s is 1000.
// Example(s): the longest palindrome in the string "ababd" is "aba" or "bab" (either is acceptable) and we would want either returned
// Inputs: asdfsda, hellollehasdf, abahellollehesdf, abacdcf

// ababd

// aba / bab

package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("abacd"))
}

// Expand at the center and find palindrome string
// Try to do that to every letter in the string
func longestPalindrome(str string) string {
	if len(str) < 2 {
		return str
	}

	longest := ""

	for i := 0; i < len(str); i++ {

		// Palindrome string's length could be odd number or even number
		// Example:
		// "aba" -> odd number
		// "abba" -> even number
		odd := expandCentre(str, i, i)
		even := expandCentre(str, i, i+1)
		long := ""

		// Get the longest string by comparing
		if len(odd) > len(even) {
			long = odd

		} else {

			long = even
		}

		if len(long) > len(longest) {

			longest = long
		}

	}

	return longest

}

// This function will find the longest string that centered with given letter
// All the indexes, positions, all status will he handled within this function
func expandCentre(s string, start, end int) string {
	for start >= 0 && end <= len(s)-1 && s[start] == s[end] {
		start -= 1
		end += 1

	}
	// After getting the longest string, start will be added 1 and end will be minus 1 anyway,
	// for slice, start need to be minues 1 and end is extactlly what we need
	// so we get this:
	return s[start+1 : end]

}
