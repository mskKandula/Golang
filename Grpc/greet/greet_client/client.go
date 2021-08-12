package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

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
	// doServerStreaming(c)

	// Client Streaming
	// doClientStreaming(c)

	// BiDi Streaming
	doBidiStreaming(c)
}

func doBidiStreaming(c greetpb.GreetClient) {

	requests := []*greetpb.GreetOnRequest{
		&greetpb.GreetOnRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mohan",
				LastName:  "Kandula",
			},
		},
		&greetpb.GreetOnRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sai",
				LastName:  "Kandula",
			},
		},
		&greetpb.GreetOnRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Krishna",
				LastName:  "Kandula",
			},
		},
		&greetpb.GreetOnRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sai",
				LastName:  "Kandula",
			},
		},
		&greetpb.GreetOnRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mohan",
				LastName:  "Kandula",
			},
		},
	}

	stream, err := c.GreetOn(context.Background())

	if err != nil {
		log.Fatalf("Error while preparing request in BiDi: %v", err)
		return
	}

	waitc := make(chan struct{})

	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while receiving the response in BiDi:%v", err)
				break
			}

			fmt.Println("The response from server is:", resp.GetResult())
		}

		close(waitc)
	}()

	for i, req := range requests {

		fmt.Printf("sending a %v request with data %v\n", i, req)

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)

		err = stream.Send(req)

		if err != nil {
			log.Fatalf("Error while sending request in BiDi: %v", err)
			break
		}
	}

	stream.CloseSend()

	<-waitc

}

func doClientStreaming(c greetpb.GreetClient) {

	requests := []*greetpb.LongGreetRequest{
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mohan",
				LastName:  "Kandula",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Sai",
				LastName:  "Kandula",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Krishna",
				LastName:  "Kandula",
			},
		},
		&greetpb.LongGreetRequest{
			Greeting: &greetpb.Greeting{
				FirstName: "Mohan",
				LastName:  "Kandula",
			},
		},
	}
	stream, err := c.LongGreet(context.Background())

	for i, req := range requests {
		stream.Send(req)
		fmt.Printf("Sending a %v request: %v\n", i, req)
		time.Sleep(1000 * time.Millisecond)
	}

	resp, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while reading a response: %v", err)
	}

	fmt.Println("The result response is :", resp.GetResult())

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
