package main

import (
	"context"
	"log"

	"./pb"

	"google.golang.org/grpc"
)

func main() {
	addr := "localhost:8080"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewWineClassifierServiceClient(conn)
	req := pb.WineReviewRequest{
		Review: "Dry with a fruity aroma",
	}

	resp, err := client.GetWineVariety(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Predicted wine variety: %v", resp.Variety)
}
