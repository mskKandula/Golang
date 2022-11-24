package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var (
	wg sync.WaitGroup
)

func SleepTime() {
	defer wg.Done()
	fmt.Println("In SleepTime Func")
	t1 := time.Now()
	time.Sleep(1 * time.Second)
	t2 := time.Since(t1)
	fmt.Println("Time taken by SleepTime :", t2)
}

func Process() {
	defer wg.Done()
	fmt.Println("In Process Func")
	t1 := time.Now()
	words(5, "")
	t2 := time.Since(t1)
	fmt.Println("Time taken by Process :", t2)
}

func words(length int, prefix string) {

	if length < 1 {
		return
	}
	letter := ""
	for i := 97; i < 123; i++ {
		letter = fmt.Sprintf("%s%c", prefix, i)
		//recursive call
		words(length-1, letter)
	}
}

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)
	go SleepTime()
	go Process()

	wg.Wait()
}
