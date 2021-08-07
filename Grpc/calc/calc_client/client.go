package main

import (
	"context"
	"fmt"
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
	// fmt.Println("Client Created", c)
}
