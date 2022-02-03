package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/tidwall/gjson"
)

const url = "https://jsonplaceholder.typicode.com/"

func main() {
	t := time.Now()
	users := FetchUsers(url)

	for _, user := range users.Array() {
		userId := user.Get("id").Int()

		posts := FetchPosts(url, int(userId))

		for _, post := range posts.Array() {
			fmt.Println(post)
		}
	}
	fmt.Println(time.Since(t))
}

func FetchPosts(url string, userId int) gjson.Result {

	postsUrl := url + "posts?userId=" + strconv.Itoa(userId)

	resp, err := http.Get(postsUrl)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	result := gjson.ParseBytes(body)

	return result

}

func FetchUsers(url string) gjson.Result {

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

	result := gjson.ParseBytes(body)

	return result

}
