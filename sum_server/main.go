package main

import (
	pb ".."
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":9090"
)

type ComputeSumServer struct {
	pb.UnimplementedSumComputerServer
}

func (s *ComputeSumServer) ComputeSum(ctx context.Context, in *pb.SumRequest) (*pb.ResultReply, error) {
	log.Printf("Received: %d , %d", in.GetFirstOperand(), in.GetSecondOperand())
	sum := in.GetSecondOperand() + in.GetFirstOperand()
	return &pb.ResultReply{Result: sum}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterSumComputerServer(s, &ComputeSumServer{})
	fmt.Printf("going to serve on port %s\n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
