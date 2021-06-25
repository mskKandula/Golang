package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello")

	http.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
