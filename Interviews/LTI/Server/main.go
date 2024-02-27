package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/mskKandula/golang/lti/pb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOperationserviceServer
}

func (s *server) PerformOperation(stream pb.Operationservice_PerformOperationServer) error {

	fmt.Println("BiDi end point hit")

	var resp int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while receiving requests in BiDi:%v", err)
			return err
		}

		num1 := req.GetFirstInteger()
		num2 := req.GetSecondInteger()
		operation := req.GetOperation()

		if operation == "+" {
			resp = num1 + num2

		} else if operation == "-" {
			resp = num1 - num2
		} else if operation == "/" {
			resp = num1 / num2
		} else {
			resp = num1 * num2
		}

		err = stream.Send(&pb.OperationResponse{
			Response: resp,
		})

		if err != nil {
			return err
		}
	}

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalln("Failed to listenL: ", err)
	}

	s := grpc.NewServer()

	// pb.RegisterOperationserviceServer(s, &server{})

	pb.RegisterOperationserviceServer(s, &server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalln("Failed to serve: ", err)
	}
}
