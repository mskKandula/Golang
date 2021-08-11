package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/mskKandula/calcpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello World")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Error dialing grpc server: ", err)
	}

	defer conn.Close()

	c := calcpb.NewCalcServiceClient(conn)

	// doUnary(c)
	// doServerStreaming(c)
	doClientStreaming(c)
}

func doUnary(c calcpb.CalcServiceClient) {
	req := &calcpb.CalcRequest{
		IntegerData: &calcpb.IntegerData{
			FirstNum:  1024,
			SecondNum: 1024,
		},
	}

	resp, err := c.Sum(context.Background(), req)

	if err != nil {
		log.Fatalln("Error fetching the response: ", err)
	}

	fmt.Println(resp)
}

func doServerStreaming(c calcpb.CalcServiceClient) {
	req := &calcpb.CalcManyRequest{
		Number: 120,
	}

	primeClient, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalln("Error fetching the response: ", err)
	}
	for {
		res, err := primeClient.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalln("Error fetching the response: ", err)
		}

		fmt.Println("The Factor is:", res.GetResult())
	}
}

func doClientStreaming(c calcpb.CalcServiceClient) {

	requests := []*calcpb.LongCalcRequest{
		&calcpb.LongCalcRequest{
			Number: 3,
		},
		&calcpb.LongCalcRequest{
			Number: 4,
		},
		&calcpb.LongCalcRequest{
			Number: 5,
		},
		&calcpb.LongCalcRequest{
			Number: 3,
		},
		&calcpb.LongCalcRequest{
			Number: 6,
		},
	}

	stream, err := c.Average(context.Background())

	if err != nil {
		log.Fatalf("Error while sending requests for client streaming")
	}

	for i, req := range requests {
		stream.Send(req)
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("Sending a %v request with data %v \n", i+1, req.GetNumber())
	}

	data, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while getting a response in Client Streaming")
	}

	fmt.Println("The result of client streaming, The Avg is: ", data.GetResult())

}
