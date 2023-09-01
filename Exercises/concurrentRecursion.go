package main

import "fmt"

func rec(n int) int {

	if n < 2 {
		return 1
	}

	ch := make(chan int)

	go func() {
		ch <- rec(n - 1)
	}()

	return rec(n-2) + <-ch

}

func main() {
	r := rec(25)

	fmt.Println(r)
}
