package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type User struct {
	Id int `json:"id"`
}

const url = "https://jsonplaceholder.typicode.com/"

func main() {
	t := time.Now()

	users := FetchUsers(url)

	for _, user := range users {
		userId := user.Id
		fmt.Println(userId)

	}

	fmt.Println(time.Since(t))
}

func FetchUsers(url string) []User {

	userUrl := url + "users"

	resp, err := http.Get(userUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var users []User

	err = json.NewDecoder(resp.Body).Decode(&users)

	if err != nil {
		log.Fatal(err)
	}

	return users

}
