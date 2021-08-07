package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mskKandula/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Error dialing grpc server: ", err)
	}

	defer conn.Close()

	c := greetpb.NewGreetClient(conn)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mohan",
			LastName:  "Kandula",
		},
	}

	resp, err := c.GreetAPI(context.Background(), req)

	if err != nil {
		log.Fatalln("Error fetching the response: ", err)
	}

	fmt.Println(resp)
	// fmt.Println("Client Created", c)
}
