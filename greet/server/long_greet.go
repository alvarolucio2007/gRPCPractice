package main

import (
	"fmt"
	"io"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
)

func (s *Server) LongGreet(stream proto.GreetService_LongGreetServer) error {
	log.Println("LongGreet function was invoked")
	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receiving %v\n", req)
		res += fmt.Sprintf("Hello %s!\n", req.FirstName)
	}
}
