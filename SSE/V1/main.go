package main

import (
	// "encoding/json"
	"log"
	"net/http"
	"strings"
	// "time"
)

func formatSSE(event string, data string) []byte {

	eventPayload := "event: " + event + "\n"
	dataLines := strings.Split(data, "\n")

	for _, line := range dataLines {
		eventPayload = eventPayload + "data: " + line + "\n"
	}

	return []byte(eventPayload + "\n")
}

func listenHandler(w http.ResponseWriter, r *http.Request) {

	_messageChannel := make(chan []byte)

	go func() {
		msgData := "Data to send"
		_messageChannel <- []byte(msgData)
		//  time.Sleep(time.Second * 3)
	}()

	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	for {
		select {

		case _msg := <-_messageChannel:
			w.Write(formatSSE("message", string(_msg)))
			w.(http.Flusher).Flush()

		case <-r.Context().Done():
			return
		}
	}
}

func main() {

	http.HandleFunc("/listen", listenHandler)

	log.Println("Running at :4000")
	log.Fatal(http.ListenAndServe(":4000", nil))
}
