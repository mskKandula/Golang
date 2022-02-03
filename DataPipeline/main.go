package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/"

func main() {
	FetchUsers(url)
}

func FetchUsers(url string) {

	userUrl := url + "users"

	resp, err := http.Get(userUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(body))

}
