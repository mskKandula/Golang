package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", pyhtonHandler)
	fmt.Println("listening on port 8081")
	http.ListenAndServe(":8081", nil)
}

func pyhtonHandler(w http.ResponseWriter, r *http.Request) {

	resp, err := http.Get("http://127.0.0.1:5000")
	if err != nil {
	  fmt.Println(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
}
fmt.Fprintf(w,string(body))

}
