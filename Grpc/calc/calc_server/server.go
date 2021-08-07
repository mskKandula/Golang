package main

import (
	"context"
	"fmt"
	"log"
	"net"

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
