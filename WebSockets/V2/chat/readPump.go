package chat

import (
	"log"

	"github.com/gorilla/websocket"
)

func (client *Client) ReadPump() {

	defer func() {
		client.Disconnect()
	}()

	for {

		_, jsonMessage, err := client.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("unexpected close error: %v", err)
			}
			break
		}
		client.wsServer.Broadcast <- jsonMessage
	}

}
