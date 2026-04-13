package main

import (
	"context"
	"log"
	"time"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *proto.GreetRequest) (*proto.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with %v\n", in)
	for range 3 {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("The client canceled the request!")
			return nil, status.Error(codes.Canceled, "client canceled request")
		}
		time.Sleep(1 * time.Second)
	}
	return &proto.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}
