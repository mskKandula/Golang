package main

import (
	"fmt"
	"time"
	"sync"
	"sync/atomic"
)
var (
	num int = 1000
	wg sync.WaitGroup
	counter uint64 = 0
)

func odd(n int,wg *sync.WaitGroup){

for i:=0; i<n; i++{

	if (i%2 !=0){

		fmt.Println(i)

		atomic.AddUint64(&counter, 1)
	}
}
defer wg.Done()
}

func main(){

	t := time.Now()

	wg.Add(1)

	go odd(num,&wg)

	for i:=0; i<num; i++{

		if(i%2 == 0){

		fmt.Println(i)

		atomic.AddUint64(&counter, 1)
	}
}
wg.Wait()

fmt.Println("Counter :",counter)

fmt.Println("Total Time Taken :",time.Since(t))
}