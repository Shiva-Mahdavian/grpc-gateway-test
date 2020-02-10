package main

import (
	"context"
	"google.golang.org/grpc"
	pb "github.com/Shiva-mahdavian/grpc-gateway-test"
	"log"
	"net"
)

const (
	port = ":50051"
)

type ComputeSumServer struct {
	pb.UnimplementedComputeSumServer
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
	pb.RegisterComputeSumServer(s, &ComputeSumServer{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
