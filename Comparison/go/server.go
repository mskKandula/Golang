package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("listening on port 9000")

	http.HandleFunc("/", requestHandler)
	log.Fatal(http.ListenAndServe(":9000", nil))
}

func fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fib := fibonacci(40) // This will take some time
	fmt.Fprintf(w, "Fibonacci result: %v", fib)
}

// autocannon -c 50 -d 10 http://localhost:9000/
