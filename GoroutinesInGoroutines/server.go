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

	go handlePayment()
	go steamMilk()
	go makeEspresso()
}

func handlePayment() {
	fmt.Println("Making Payment")
	time.Sleep(2 * time.Second)
}

func steamMilk() {
	fmt.Println("Steaming Milk")
	time.Sleep(2 * time.Second)
}

func makeEspresso() {
	fmt.Println("Making Espresso")
	time.Sleep(2 * time.Second)
}
