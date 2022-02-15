package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
	"time"
)

var wg sync.WaitGroup

type User struct {
	Id int `json:"id"`
}
type Album struct {
	Id int `json:"id"`
}

type Photo struct {
	AlbumId int    `json:"albumId"`
	Url     string `json:"url"`
}

type Image struct {
	ImageData []byte
	FileName  string
	Extension string
}

// file path creation
func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0770); err != nil {
		return nil, err
	}
	return os.Create(p)
}

const url = "https://jsonplaceholder.typicode.com/"

func main() {
	t := time.Now()

	userIdsChan := make(chan int, 10)
	albumIdsChan := make(chan []Album, 10)
	albumIdChan := make(chan int, 100)
	photosChan := make(chan []Photo, 50)
	photoChan := make(chan Photo, 1000)
	resultChan := make(chan Image, 1000)
	boolChan := make(chan bool)

	go FetchAlbums(userIdsChan, albumIdsChan)
	go ProcessAlbums(albumIdsChan, albumIdChan)
	go FetchPhotos(albumIdChan, photosChan)
	go ProcessPhotos(photosChan, photoChan)

	for i := 0; i < 500; i++ {
		go FetchImages(photoChan, resultChan, &wg)
		wg.Add(1)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	go SaveImages(resultChan, boolChan)

	users := FetchUsers(url)

	for _, user := range users {
		userId := user.Id
		userIdsChan <- userId

	}
	close(userIdsChan)
	<-boolChan

	fmt.Println("Total time taken : ", time.Since(t))
}

func SaveImages(resultChan <-chan Image, boolChan chan<- bool) {
	err := os.Chdir("Images")

	if err != nil {
		log.Fatal(err)
	}

	for image := range resultChan {

		fileName := image.FileName + "." + image.Extension

		file, err := create(fileName)

		if err != nil {
			log.Fatal(err)
		}

		io.Copy(file, bytes.NewReader(image.ImageData))

		file.Close()

	}
	boolChan <- true

}

func FetchImages(photoChan <-chan Photo, resultChan chan<- Image, wg *sync.WaitGroup) {
	for photo := range photoChan {

		url := photo.Url

		resp, err := http.Get(url)

		if err != nil {
			log.Fatal(err)
		}

		arr := strings.Split(url, "/")

		fileName := strconv.Itoa(photo.AlbumId) + "/" + arr[len(arr)-1]

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
	wg.Done()
}

func ProcessPhotos(photosChan <-chan []Photo, photoChan chan<- Photo) {
	for photos := range photosChan {
		for _, photo := range photos {
			photoChan <- photo
		}
	}
	close(photoChan)
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
