package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mskKandula/calcpb"
	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *calcpb.CalcRequest) (*calcpb.CalcResponse, error) {
	firstNum := req.GetIntegerData().GetFirstNum()
	lastNum := req.GetIntegerData().GetFirstNum()

	result := int64(firstNum + lastNum)

	res := &calcpb.CalcResponse{
		Result: result,
	}

	return res, nil

}

func (*server) Prime(req *calcpb.CalcManyRequest, stream calcpb.CalcService_PrimeServer) error {
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
