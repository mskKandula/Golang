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

	postsChan := make(chan int)
	processPostsChan := make(chan gjson.Result)
	commentsChan := make(chan int)
	processCommentsChan := make(chan gjson.Result)
	resultChan := make(chan gjson.Result)
	boolChan := make(chan bool)
	go FetchPosts(postsChan, processPostsChan)
	go ProcessPosts(processPostsChan, commentsChan)
	go FetchComments(commentsChan, processCommentsChan)
	go ProcessComments(processCommentsChan, resultChan)
	go PrintResult(resultChan, boolChan)

	t := time.Now()
	users := FetchUsers(url)

	for _, user := range users.Array() {
		userId := user.Get("id").Int()
		postsChan <- int(userId)

	}
	close(postsChan)
	<-boolChan
	fmt.Println(time.Since(t))
}

func PrintResult(resultChan <-chan gjson.Result, boolChan chan<- bool) {
	for results := range resultChan {
		fmt.Println(results)

	}
	boolChan <- true
}

func ProcessComments(processCommentsChan <-chan gjson.Result, resultChan chan<- gjson.Result) {
	for comments := range processCommentsChan {

		for _, comment := range comments.Array() {

			resultChan <- comment
		}
	}
	close(resultChan)
}

func FetchComments(commentsChan <-chan int, processCommentsChan chan<- gjson.Result) {

	for postId := range commentsChan {

		userUrl := url + "comments?postId=" + strconv.Itoa(postId)

		resp, err := http.Get(userUrl)

		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		processCommentsChan <- gjson.ParseBytes(body)
		resp.Body.Close()
	}

	close(processCommentsChan)

}

func ProcessPosts(processPostsChan <-chan gjson.Result, commentsChan chan<- int) {
	for posts := range processPostsChan {

		for _, post := range posts.Array() {

			postId := post.Get("id").Int()
			commentsChan <- int(postId)
		}
	}
	close(commentsChan)
}

func FetchPosts(postsChan <-chan int, processPostsChan chan<- gjson.Result) {

	for userId := range postsChan {

		postsUrl := url + "posts?userId=" + strconv.Itoa(userId)

		resp, err := http.Get(postsUrl)

		if err != nil {
			log.Fatal(err)
		}

		body, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		processPostsChan <- gjson.ParseBytes(body)
		resp.Body.Close()
	}
	close(processPostsChan)
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
