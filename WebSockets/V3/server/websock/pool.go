package websock

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/mailru/easygo/netpoll"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[string][]*Client
	Broadcast  chan []byte
}

var (
	poolInit       *Pool
	ClientConnChan chan *Client
)

func init() {
	poolInit = &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[string][]*Client),
		Broadcast:  make(chan []byte),
	}
	ClientConnChan = make(chan *Client, 200)
}

func NewPool() *Pool {
	return poolInit
}

func (pool *Pool) Start(db *sql.DB) {

	poller, err := netpoll.New(nil)
	if err != nil {
		log.Println(err)
	}

	stmt, err := db.Prepare("INSERT INTO Messages(title) VALUES(?)")
	if err != nil {
		log.Println(err)
	}

	for {

		select {

		case client := <-pool.Register:

			pool.Clients["test"] = append(pool.Clients["test"], client)

			// Get netpoll descriptor with EventRead|EventEdgeTriggered.
			desc := netpoll.Must(netpoll.HandleRead(client.Conn))

			// Make conn to be observed by netpoll instance.
			poller.Start(desc, func(ev netpoll.Event) {
				if ev&netpoll.EventReadHup != 0 {
					poller.Stop(desc)
					client.Conn.Close()
					return
				}
				ClientConnChan <- client
			})

			if len(pool.Clients["test"])%1000 == 0 {
				fmt.Println("Size of Connection Pool: ", len(pool.Clients["test"]))
			}

		case <-pool.Unregister:

			delete(pool.Clients, "test")

			fmt.Println("Size of Connection Pool : ", len(pool.Clients["test"]))

		case msg := <-pool.Broadcast:
			_, err = stmt.Exec(msg)
			if err != nil {
				log.Println(err)
			}
		}
	}
}
