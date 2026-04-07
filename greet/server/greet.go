package main

import (
	"context"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &proto.GreetResponse{
		Result: "Hello: " + in.FirstName,
	}, nil
}
