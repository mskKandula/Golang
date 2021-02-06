package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"time"
)

// listen for a broadcast (effecient)
func listen(name string, a map[string]int, c *sync.Cond) {
	c.L.Lock()
	c.Wait() //waiting for broadcast
	fmt.Println(name, " age:", a["T"])
	c.L.Unlock()
}

func broadcast(name string, a map[string]int, c *sync.Cond) {
	time.Sleep(time.Second)
	c.L.Lock()
	a["T"] = 25
	c.Broadcast() //Broadcasting to all the goroutines which are waiting
	c.L.Unlock()
}

func main() {

	var age = make(map[string]int) //using map since it's a reference type

	m := sync.Mutex{}
	cond := sync.NewCond(&m)

	// listener 1
	go listen("lis1", age, cond)

	// listener 2
	go listen("lis2", age, cond)

	// broadcast
	go broadcast("b1", age, cond)

	//you can use waitgroups also from sync package
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt) //run till any interruption occurs
	<-ch
}
