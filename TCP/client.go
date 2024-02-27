package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("abc"))
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)

	_, err = conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(buffer))

}
