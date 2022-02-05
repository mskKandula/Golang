package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type User struct {
	Id int `json:"id"`
}
type Album struct {
	Id int `json:"id"`
}

type Photos struct {
	Url string `json:"url"`
}

const url = "https://jsonplaceholder.typicode.com/"

func main() {
	t := time.Now()

	userIdsChan := make(chan int)
	albumIdsChan := make(chan []Album)
	albumIdChan := make(chan int)
	boolChan := make(chan bool)

	go FetchAlbums(userIdsChan, albumIdsChan)
	go ProcessAlbums(albumIdsChan, albumIdChan)

	go func(albumIdChan <-chan int, boolChan chan<- bool) {

		for id := range albumIdChan {
			fmt.Println(id)
		}
		boolChan <- true

	}(albumIdChan, boolChan)

	users := FetchUsers(url)

	for _, user := range users {
		userId := user.Id
		userIdsChan <- userId

	}
	close(userIdsChan)
	<-boolChan
	fmt.Println(time.Since(t))
}

func ProcessAlbums(albumIdsChan <-chan []Album, albumIdChan chan<- int) {
	for albums := range albumIdsChan {
		for _, album := range albums {
			albumIdChan <- album.Id
		}
	}
	close(albumIdChan)
}

func FetchAlbums(userIdsChan <-chan int, albumIdsChan chan<- []Album) {

	for id := range userIdsChan {

		albumUrl := url + "albums?userId=" + strconv.Itoa(id)

		resp, err := http.Get(albumUrl)

		if err != nil {
			log.Fatal(err)
		}

		var albums []Album

		err = json.NewDecoder(resp.Body).Decode(&albums)

		resp.Body.Close()

		if err != nil {
			log.Fatal(err)
		}

		albumIdsChan <- albums
	}

	close(albumIdsChan)

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
