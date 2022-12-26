package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
	"time"
)

func main() {
	http.HandleFunc("/coffeShop", coffeShop)

	fmt.Println("listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func coffeShop(w http.ResponseWriter, r *http.Request) {
	var (
		i  int
		wg sync.WaitGroup
	)

	t1 := time.Now()
	id := r.URL.Query().Get("id")
	val, _ := strconv.Atoi(id)

	for i < val {
		wg.Add(1)
		go makeCoffee(&wg)
		i++
	}

	wg.Wait()
	fmt.Fprintf(w, "Time taken to serve: %v", time.Since(t1))
}

func makeCoffee(wg *sync.WaitGroup) {
	defer wg.Done()

	var newWg sync.WaitGroup
	newWg.Add(3)

	go handlePayment(&newWg)
	go steamMilk(&newWg)
	go makeEspresso(&newWg)

	newWg.Wait()
}

func handlePayment(newWg *sync.WaitGroup) {
	defer newWg.Done()
	fmt.Println("Making Payment")
	time.Sleep(2 * time.Second)
}

func steamMilk(newWg *sync.WaitGroup) {
	defer newWg.Done()
	fmt.Println("Steaming Milk")
	time.Sleep(2 * time.Second)
}

func makeEspresso(newWg *sync.WaitGroup) {
	defer newWg.Done()
	fmt.Println("Making Espresso")
	time.Sleep(2 * time.Second)
}
