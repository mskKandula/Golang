package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	http.HandleFunc("/coffeShop", coffeShop)

	fmt.Println("listening on port 9000")
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func coffeShop(w http.ResponseWriter, r *http.Request) {
	var (
		i int
	)
	t1 := time.Now()
	id := r.URL.Query().Get("id")
	val, _ := strconv.Atoi(id)
	for i < val {
		makeCoffee()
		i++
	}

	fmt.Fprintf(w, "Time taken to serve: %v", time.Since(t1))
}

func makeCoffee() {

	handlePayment()
	steamMilk()
	makeEspresso()
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
