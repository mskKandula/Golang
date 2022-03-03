package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type broadcastMsg struct {
	Message map[string]interface{}
	RoomId  string
	Client  *websocket.Conn
}

var (
	room     Room
	upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	broadcast = make(chan broadcastMsg)
)

func CreateRoom(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")

	roomID := room.RoomCreation()

	json.NewEncoder(w).Encode(map[string]string{
		"roomId": roomID,
	})
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	roomId, ok := r.URL.Query()["roomId"]
	if !ok {
		log.Println("roomId missing in URL Parameters")
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}

	room.InsertIntoRoom(roomId[0], ws, false)

	go broadcaster()

	for {
		var msg broadcastMsg

		err := ws.ReadJSON(&msg.Message)

		if err != nil {
			log.Fatal("Read Error: ", err)
		}

		msg.Client = ws
		msg.RoomId = roomId[0]

		log.Println(msg.Message)

		broadcast <- msg
	}
}

func broadcaster() {
	for {
		msg := <-broadcast

		for _, client := range room.Users[msg.RoomId] {
			if client.Conn != msg.Client {
				err := client.Conn.WriteJSON(msg.Message)

				if err != nil {
					log.Fatal(err)
					client.Conn.Close()
					delete(room.Users, msg.RoomId)
				}
			}
		}
	}
}
