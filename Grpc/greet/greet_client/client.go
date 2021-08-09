package main

import (
	"context"
	"fmt"
	"io"
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
	// Unary
	// doUnary(c)

	// Server Streaming
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetClient) {
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
}

func doServerStreaming(c greetpb.GreetClient) {

	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Mohana",
			LastName:  "Kandula",
		},
	}

	respStream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalln("Error fetching  the streaming response: ", err)
	}
	for {
		msg, err := respStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error while reading the stream:", err)
		}

		fmt.Println("Response from streaming is : ", msg.GetResult())
	}
}
