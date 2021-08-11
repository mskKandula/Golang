package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/mskKandula/greetpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) GreetAPI(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	result := "Hello" + " " + firstName + " " + lastName

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil

}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.Greet_GreetManyTimesServer) error {

	fmt.Println("Server streaming")

	firstName := req.GetGreeting().GetFirstName()
	lastName := req.GetGreeting().GetLastName()

	for i := 0; i < 10; i++ {
		result := "Hello" + " " + firstName + " " + lastName + " Number :" + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}

		stream.Send(res)

		time.Sleep(1000 * time.Millisecond)
	}

	return nil

}

func (*server) LongGreet(stream greetpb.Greet_LongGreetServer) error {
	fmt.Println("Long Greet End Point Hit")

	var response string

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(
				&greetpb.LongGreetResponse{
					Result: response,
				},
			)
		}

		if err != nil {
			log.Fatalf("Error while reading the client stream: %v", err)
		}

		firstName := msg.GetGreeting().GetFirstName()

		response += "\nHello " + firstName
	}
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalln("Failed to listenL: ", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreetServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}
