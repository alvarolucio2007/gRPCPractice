package main

import (
	"context"
	"fmt"
	"log"
	"math"

	"github.com/alvarolucio2007/gRPCPractice/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *proto.SqrtRequest) (*proto.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with %v\n", in)
	number := in.Number
	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Received a negative number: %d", number),
		)
	}
	return &proto.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
