package main

import (
	"fmt"
	"time"
)

func main(){
	startTime :=time.Now()

	jobs :=make(chan int,100)

	results:=make(chan int,100)

	NoWorkers:= 10

	for k:=0;k<NoWorkers;k++{

		go worker(jobs,results)

	}

	for i:=0;i<100;i++{
		jobs <- i
	}

	close(jobs)

	for j:= range results{
		fmt.Println(j)
	}
	fmt.Printf("Total Time taken  is : %v ",time.Since(startTime))

}

func worker(jobs <-chan int,results chan<- int){

	for n := range jobs{

		results <- fib(n)

	}
	close(results)
}

func fib(n int)int{

	if n<=1{
		return 1
	}

	return fib(n-1) + fib(n-2)
}
