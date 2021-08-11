package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"github.com/mskKandula/calcpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {

	fmt.Println("Unary End Point hit")

	firstNum := req.GetIntegerData().GetFirstNum()
	lastNum := req.GetIntegerData().GetFirstNum()

	result := int64(firstNum + lastNum)

	res := &calcpb.CalcResponse{
		Result: result,
	}

	return res, nil

}

func (*server) Prime(req *calcpb.CalcManyRequest, stream calcpb.CalcService_PrimeServer) error {

	fmt.Println("Server Streaming End Point hit")

	N := req.GetNumber()
	var k int32 = 2

	for N > 1 {
		if N%k == 0 { // if k evenly divides into N
			time.Sleep(1000 * time.Millisecond)
			resp := &calcpb.CalcManyResponse{
				Result: k, // this is a factor
			}
			stream.Send(resp)
			N = N / k // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}
	}
	return nil
}

func (*server) Average(stream calcpb.CalcService_AverageServer) error {

	fmt.Println("Client Streaming End Point hit")

	var (
		i     int32   = 0
		total int32   = 0
		avg   float32 = 0
	)
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(
				&calcpb.LongCalcResponse{
					Result: avg,
				},
			)
		}

		if err != nil {
			log.Fatalf("Error while receiving requests: %v", err)
		}

		number := req.GetNumber()

		i++
		total += number
		avg = (float32(total) / float32(i))
	}
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalln("Failed to listenL: ", err)
	}

	s := grpc.NewServer()

	calcpb.RegisterCalcServiceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}
