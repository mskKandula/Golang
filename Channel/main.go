package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	// ch2:=make(chan int,1)
	// ch3:= make(chan int,2)
	ch1 <- 1
	fmt.Println(<-ch1)

	ch1 <- 2
	fmt.Println(<-ch1)

	// time.Sleep(1 * time.Second)

	// ch2 <-2
	// ch3 <-3

}
