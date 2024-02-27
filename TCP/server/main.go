package main

import (
	"fmt"
	"log"
	"net"
)

func main() {

	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	buffer := make([]byte, 1024)

	_, err := conn.Read(buffer)
	if err != nil {
		conn.Close()
		log.Println(err)
	}

	str := string(buffer)
	fmt.Println("Incoming Message:", str)

	result := ""
	for _, s := range str {
		result = string(s) + result
	}

	conn.Write([]byte(result))

	// reader := bufio.NewReader(conn)

	// for {
	// 	msg, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		log.Println(err)
	// 		conn.Close()
	// 		return
	// 	}

	// 	fmt.Println("Incoming Message:", msg)
	// 	conn.Write([]byte("Message Recieved.\n"))
	// }

}
