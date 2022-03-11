package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

var (
	counter     int
	clients     map[string]*Client
	clientsList []string
)

type (
	Client struct {
		name   string
		events chan *DashBoard
	}
	Currency float64
	Item     struct {
		Name     string   `json:"name,omitempty"`
		Quantity int      `json:"quantity,omitempty"`
		Price    Currency `json:"price,omitempty"`
	}
	Store struct {
		Items map[string]Item `json:"items,omitempty"`
	}
	DashBoard struct {
		Users         uint       `json:"users,omitempty"`
		UsersLoggedIn uint       `json:"users_logged_in,omitempty"`
		Inventory     *Store     `json:"inventory,omitempty"`
		ChartOne      []int      `json:"chart_one,omitempty"`
		ChartTwo      []Currency `json:"chart_two,omitempty"`
	}
)

func main() {
	fmt.Println("Hello")
	clients = make(map[string]*Client)
	go updateDashboard()
	// register static files handle '/index.html -> client/index.html'
	// http.Handle("/", http.FileServer(http.Dir("client")))
	// register RESTful handler for '/sse/dashboard'
	http.HandleFunc("/sse/dashboard", dashbaordHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dashbaordHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("Client: %v", r.RemoteAddr)
	client := clients[r.RemoteAddr]
	if nil == client {
		client = addClient(r.RemoteAddr)
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	timeout := time.After(1 * time.Second)
	select {
	case ev := <-client.events:
		var buf bytes.Buffer
		enc := json.NewEncoder(&buf)
		enc.Encode(ev)
		fmt.Fprintf(w, "data: %v\n\n", buf.String())
		fmt.Printf("data: %v\n", buf.String())
	case <-timeout:
		fmt.Fprintf(w, ": nothing to sent\n\n")
	}

	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}

func updateDashboard() {
	for {
		inv := updateInventory()
		db := &DashBoard{
			Users:         uint(rand.Uint32()),
			UsersLoggedIn: uint(rand.Uint32() % 200),
			Inventory:     inv,
			ChartOne:      []int{2, 35, 634, 93, 390432},
			ChartTwo:      []Currency{3.59, 6.09, 563.90},
		}

		client := getClient()
		if nil != client {
			client.events <- db
		}
	}
}

func addClient(s string) *Client {
	c := &Client{name: s, events: make(chan *DashBoard, 10)}
	clients[s] = c
	clientsList = append(clientsList, s)
	return c
}

func getClient() *Client {
	if 0 == len(clientsList) {
		return nil
	}

	r := rand.Int() % len(clientsList)
	s := clientsList[r]
	return clients[s]
}
