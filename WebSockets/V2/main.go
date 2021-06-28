package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mskKandula/chat"
)

func main() {
	fmt.Println("Hello")

	wsServer := chat.NewWebsocketServer()

	go wsServer.Run()

	http.HandleFunc("/", homePage)

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chat.ServeWs(wsServer, w, r)
	})

	log.Fatal(http.ListenAndServe(":8090", nil))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}
