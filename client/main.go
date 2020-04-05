package main

import (
	"context"
	"log"
	"time"

	pb "../sum"
	"google.golang.org/grpc"
)

const (
	address = "localhost:9090"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewSumComputerClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.ComputeSum(ctx, &pb.SumRequest{FirstOperand: 23, SecondOperand: 12})
	if err != nil {
		log.Fatalf("could send sum request: %v", err)
	}
	log.Printf("sum result: %d", r.GetResult())
}
