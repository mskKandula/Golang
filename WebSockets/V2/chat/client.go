package chat

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  4096,
	WriteBufferSize: 4096,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

// Client represents the websocket client at the server
type Client struct {
	// The actual websocket connection.
	Conn     *websocket.Conn
	WsServer *WsServer
	Rooms map[*Room]bool
}

func newClient(conn *websocket.Conn, wsServer *WsServer) *Client {
	return &Client{
		Conn:     conn,
		WsServer: wsServer,
		Rooms: make(map[*Room]bool)
	}
}

func (client *Client) Disconnect() {
    client.wsServer.UnRegister <- client
    for room := range client.Rooms {
        room.UnRegister <- client
    }
}

// ServeWs handles websocket requests from clients requests.
func ServeWs(wsServer *WsServer, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := newClient(conn, wsServer)

	go client.ReadPump()

	fmt.Println("New Client Joined")

	fmt.Println("client : ", client)

	wsServer.Register <- client
}
