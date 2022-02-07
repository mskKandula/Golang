package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type User struct {
	Id int `json:"id"`
}
type Album struct {
	Id int `json:"id"`
}

type Photo struct {
	Url string `json:"url"`
}

type Image struct {
	ImageData []byte
	FileName  string
	Extension string
}

const url = "https://jsonplaceholder.typicode.com/"

func main() {
	t := time.Now()

	userIdsChan := make(chan int)
	albumIdsChan := make(chan []Album)
	albumIdChan := make(chan int)
	photosChan := make(chan []Photo)
	urlChan := make(chan string)
	resultChan := make(chan Image)
	boolChan := make(chan bool)

	go FetchAlbums(userIdsChan, albumIdsChan)
	go ProcessAlbums(albumIdsChan, albumIdChan)
	go FetchPhotos(albumIdChan, photosChan)
	go ProcessPhotos(photosChan, urlChan)
	go FetchImages(urlChan, resultChan)

	go func(resultChan <-chan Image, boolChan chan<- bool) {
		err := os.Chdir("Images")

		if err != nil {
			log.Fatal(err)
		}

		for image := range resultChan {

			fileName := image.FileName + "." + image.Extension

			file, err := os.Create(fileName)

			if err != nil {
				log.Fatal(err)
			}

			io.Copy(file, bytes.NewReader(image.ImageData))

			file.Close()

		}
		boolChan <- true

	}(resultChan, boolChan)

	users := FetchUsers(url)

	for _, user := range users {
		userId := user.Id
		userIdsChan <- userId

	}
	close(userIdsChan)
	<-boolChan
	fmt.Println(time.Since(t))
}

func FetchImages(urlChan <-chan string, resultChan chan<- Image) {
	for url := range urlChan {

		resp, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		arr := strings.Split(url, "/")
		fileName := arr[len(arr)-1]

		contentType := resp.Header.Get("Content-Type")
		extension := strings.Split(contentType, "/")[1]

		// contentLength := resp.ContentLength

		imageData, err := io.ReadAll(resp.Body)

		if err != nil {
			log.Fatal(err)
		}

		resp.Body.Close()

		resultChan <- Image{
			imageData,
			fileName,
			extension,
		}
	}
	close(resultChan)
}

func ProcessPhotos(photosChan <-chan []Photo, urlChan chan<- string) {
	for photos := range photosChan {
		for _, photo := range photos {
			urlChan <- photo.Url
		}
	}
	close(urlChan)
}

func FetchPhotos(albumIdChan <-chan int, photosChan chan<- []Photo) {
	for id := range albumIdChan {

		photosUrl := url + "photos?albumId=" + strconv.Itoa(id)

		resp, err := http.Get(photosUrl)

		if err != nil {
			log.Fatal(err)
		}

		var photos []Photo

		err = json.NewDecoder(resp.Body).Decode(&photos)

		resp.Body.Close()

		if err != nil {
			log.Fatal(err)
		}

		photosChan <- photos
	}
	close(photosChan)
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
