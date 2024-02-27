package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	"github.com/mskKandula/golang/lti/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalln("Error dialing grpc server: ", err)
	}

	defer conn.Close()

	c := pb.NewOperationserviceClient(conn)

	doBiDiStreaming(c)
}

func doBiDiStreaming(c pb.OperationserviceClient) {

	request1 := []int32{10, 9, 8, 7}

	request2 := []int32{9, 8, 7, 6}

	request3 := []string{"+", "-", "/", "*"}

	waitc := make(chan struct{})

	stream, err := c.PerformOperation(context.Background())

	if err != nil {
		log.Fatalf("Error %v", err)
		return
	}

	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error while recieving responses in BiDi:%v", err)
				break
			}
			fmt.Println("The Result is:", resp.GetResponse())
		}
		close(waitc)
	}()

	for i := range request1 {
		err = stream.Send(&pb.OperationRequest{
			FirstInteger:  request1[i],
			SecondInteger: request2[i],
			Operation:     request3[i],
		})

		if err != nil {
			log.Fatalf("Error while sending request in BiDi: %v", err)
			break
		}

		time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
	}

	stream.CloseSend()

	<-waitc
}
