package websock

import (
	"fmt"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string][]*Client
	Broadcast  chan []byte
}

var poolInit *Pool

func init() {
	poolInit = &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string][]*Client),
		Broadcast:  make(chan []byte),
	}
}

func NewPool() *Pool {
	return poolInit
}

func (pool *Pool) Start() {

	for {

		select {

		case client := <-pool.Register:

			pool.Clients["test"] = append(pool.Clients["test"], client)

			if len(pool.Clients["test"])%1000 == 0 {
				fmt.Println("Size of Connection Pool: ", len(pool.Clients["test"]))
			}

		case <-pool.Unregister:

			delete(pool.Clients, "test")

			fmt.Println("Size of Connection Pool : ", len(pool.Clients["test"]))

		case msg := <-pool.Broadcast:
			fmt.Println(string(msg))
		}
	}
}
