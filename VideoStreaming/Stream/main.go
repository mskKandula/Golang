package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	fmt.Println("Hello")

	http.HandleFunc("/", homePage)

	http.HandleFunc("/media/{mId:[0-9]+}", streamHandler)

	http.HandleFunc("/media/{mId:[0-9]+}/{segName:index[0-9]+.ts}", streamHandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}
func streamHandler(w http.ResponseWriter, r *http.Request) {

	mId, err := strconv.Atoi(r.URL.Query().Get("mId"))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	segName := r.URL.Query().Get("segName")
	fmt.Println("out")
	if segName == "" {
		fmt.Println("In it")
		mediaBase := getMediaBase(mId)
		fmt.Println(mediaBase)
		m3u8Name := "index.m3u8"
		serveHlsM3u8(w, r, mediaBase, m3u8Name)
	} else {
		mediaBase := getMediaBase(mId)
		serveHlsTs(w, r, mediaBase, segName)
	}
}

func getMediaBase(mId int) string {
	mediaRoot := "../assets/media"
	return fmt.Sprintf("%s/%d", mediaRoot, mId)
}

func serveHlsM3u8(w http.ResponseWriter, r *http.Request, mediaBase, m3u8Name string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, m3u8Name)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "application/x-mpegURL")
}

func serveHlsTs(w http.ResponseWriter, r *http.Request, mediaBase, segName string) {
	mediaFile := fmt.Sprintf("%s/%s", mediaBase, segName)
	http.ServeFile(w, r, mediaFile)
	w.Header().Set("Content-Type", "video/MP2T")
}
