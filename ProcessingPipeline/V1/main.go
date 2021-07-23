package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.Now()

	for i := 0; i < 7; i++ {
		fmt.Println(stage3(stage2(stage1(i))))
	}

	fmt.Println(time.Since(t1))
}

func stage1(i int) int {

	time.Sleep(1 * time.Second)
	return i

}

func stage2(j int) int {

	time.Sleep(1 * time.Second)
	return j

}

func stage3(k int) int {

	time.Sleep(3 * time.Second)
	return k

}
