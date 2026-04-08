package main

import (
	"fmt"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
)

func (s *Server) GreetManyTimes(in *proto.GreetRequest, stream proto.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v\n", in)
	for i := range 10 {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i)
		if err := stream.Send(&proto.GreetResponse{
			Result: res,
		}); err != nil {
			return err
		}
	}
	return nil
}
