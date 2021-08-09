package main

import (
	"context"
	"fmt"
	"io"
	"log"

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
	doServerStreaming(c)
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
