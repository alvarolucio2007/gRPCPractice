package main

import (
	"context"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *proto.SumRequest) (*proto.SumResponse, error) {
	log.Printf("Sum function was invoked with %v\n", in)
	sum := in.FirstNumber + in.SecondNumber
	return &proto.SumResponse{
		Result: sum,
	}, nil
}
