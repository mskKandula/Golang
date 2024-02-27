package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func main() {
	r := gin.Default()

	pprof.Register(r)

	r.GET("/ws", serveWs)
	r.Run(":9000")

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func serveWs(c *gin.Context) {

	// upgrade this connection to a WebSocket connection
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// writes new messages indefinitely to our WebSocket connection
	reader(ws)
}

func reader(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()

	for {

		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("msg: %s, msgType: %d", string(msg), msgType)
	}

}
