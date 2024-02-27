package main

func main() {
	strArr := []string{"abc", "ab", "abcde", "a", "abcd"}
	n := len(strArr)
	j := n - 1
	for i := 0; i < n; i++ {
		if len(strArr[i])%2 == 0 && len(strArr[j])%2 == 0 {

		}
	}
}
