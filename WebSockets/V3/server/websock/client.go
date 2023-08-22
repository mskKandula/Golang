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

func (c *Client) Read() {
	defer func() {
		c.Pool.Unregister <- c
		c.Conn.Close()
	}()

	// Start endless read loop, waiting for messages from client
	for {

		byteData, _, err := wsutil.ReadClientData(c.Conn)
		if err != nil {
			log.Println(err)
			break
		}

		c.Pool.Broadcast <- byteData

		fmt.Printf("Message Received: %+v\n", string(byteData))
	}
}
