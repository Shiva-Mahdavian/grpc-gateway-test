package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/Shiva-mahdavian/grpc-gateway-test"
)

const (
	address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewComputeSumClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.ComputeSum(ctx, &pb.SumRequest{FirstOperand: 23, SecondOperand: 12})
	if err != nil {
		log.Fatalf("could send sum request: %v", err)
	}
	log.Printf("sum result: %d", r.GetResult())
}
