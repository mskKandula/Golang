package main

import (
	"fmt"
	"net/http"

	"golang.org/x/sync/singleflight"
)

var g = singleflight.Group{}

func main() {
	http.HandleFunc("/api/v1/get_something", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		response, _, _ := g.Do(name, func() (interface{}, error) {
			result := processingRequest(name)
			return result, nil
		})

		fmt.Fprint(w, response)
	})

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}

func processingRequest(name string) string {
	fmt.Println("[DEBUG] processing request..")
	return "Hi there! You requested " + name
}
