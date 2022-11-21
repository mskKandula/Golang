package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	wg   sync.WaitGroup
	N    int = 5
	Eats int = 3
)

func sleepRand() {
	r := rand.Int31n(10) + 1
	time.Sleep(time.Millisecond * time.Duration(r))
}

func takeLeft(num int, leftFork *sync.Mutex) {
	sleepRand()
	leftFork.Lock()
	fmt.Printf("Philosopher %v has taken the left fork %v\n", num, num)
}

func takeRight(num int, rightFork *sync.Mutex) {
	sleepRand()
	rightFork.Lock()
	fmt.Printf("Philosopher %v has taken the right fork %v\n", num, (num+1)%N)
}

func eatRelease(num int, leftFork, rightFork *sync.Mutex) {
	sleepRand()
	fmt.Printf("========Philosopher %d has done eating=========\n", num)
	leftFork.Unlock()
	rightFork.Unlock()
	fmt.Printf("Philosopher %d has done eating & released forks\n", num)
}

func philo(num int, forks *[5]sync.Mutex) {
	defer wg.Done()
	for i := 0; i < Eats; i++ {
		left := num
		right := (num + 1) % N
		if left < right {
			takeLeft(num, &forks[left])
			takeRight(num, &forks[right])
			eatRelease(num, &forks[left], &forks[right])
		} else {
			takeRight(num, &forks[right])
			takeLeft(num, &forks[left])
			eatRelease(num, &forks[left], &forks[right])
		}
	}

}

func main() {
	fmt.Printf("%v philosophers begin to eat\n", N)

	forks := [5]sync.Mutex{}

	for i := 0; i < N; i++ {
		wg.Add(1)
		go philo(i, &forks)
	}
	wg.Wait()
}
