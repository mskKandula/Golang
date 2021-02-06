package main

import (
	"fmt"
	"sync"
	"time"
)

//listens continuously till map data changes (utilizes resources)
func listen(name string, a map[string]int, wg *sync.WaitGroup) {

	for {
		if a["T"] != 0 {

			fmt.Println(name, " age:", a["T"])

			break
		}
	}
	wg.Done()
}

func broadcast(name string, a map[string]int) {
	time.Sleep(time.Second * 3)
	a["T"] = 25

}

func main() {

	var age = make(map[string]int)

	var wg sync.WaitGroup

	go listen("lis1", age, &wg)
	wg.Add(1)

	// listener 2
	go listen("lis2", age, &wg)
	wg.Add(1)

	// broadcast
	go broadcast("b1", age)

	wg.Wait()
}
