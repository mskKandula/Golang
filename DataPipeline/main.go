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
			postId := post.Get("id").Int()

			comments := FetchComments(url, int(postId))

			for _, comment := range comments.Array() {
				fmt.Println(comment)
			}
		}
	}
	fmt.Println(time.Since(t))
}

func FetchComments(url string, postId int) gjson.Result {

	userUrl := url + "comments?postId=" + strconv.Itoa(postId)

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
