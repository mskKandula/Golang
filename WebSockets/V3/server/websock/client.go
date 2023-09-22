package websock

import (
	"fmt"
	"log"
	"net"

	"github.com/gobwas/ws/wsutil"
)

type Client struct {
	Conn net.Conn
	Pool *Pool
}

func Read(clients <-chan *Client) {
	// defer func() {
	// 	c.Pool.Unregister <- c
	// 	c.Conn.Close()
	// }()

	// Start endless read loop, waiting for messages from client
	// for {
	for c := range clients {

		byteData, _, err := wsutil.ReadClientData(c.Conn)
		if err != nil {
			log.Println(err)
			return
		}

		c.Pool.Broadcast <- byteData

		fmt.Printf("Message Received: %+v\n", string(byteData))
		// }
	}
}
