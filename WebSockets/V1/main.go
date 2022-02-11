package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type IntegerArray struct {
	Intarr []int `json : intarr`
}

// We'll need to define an Upgrader
// this will require a Read and Write buffer size
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// define a writer which will write new messages sent to our WebSocket connection
func writer(conn *websocket.Conn) {

	for {
		rand.Seed(time.Now().UnixNano())

		i := &IntegerArray{

			// creating an integer array of size 12 with random elements
			Intarr: rand.Perm(12),
		}

		if err := conn.WriteJSON(i); err != nil {
			log.Println(err)
			return
		}

		time.Sleep(3 * time.Second)
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {

	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected")

	err = ws.WriteMessage(1, []byte("Hi Client!"))
	if err != nil {
		log.Println(err)
		return
	}

	// writes new messages indefinitely to our WebSocket connection
	writer(ws)
}

func main() {
	fmt.Println("Hello World")

	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", wsEndpoint)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
