package main

import (
	"fmt"
	"net/http"
)

func main() {

	url := "https://via.placeholder.com/600/92c952"
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Content Type :", resp.Header.Get("Content-Type"))
	fmt.Printf("Content Length %d bytes", resp.ContentLength)
}
