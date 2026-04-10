package main

import (
	"io"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/max/proto"
)

func (s *Server) Max(stream proto.MaxService_MaxServer) error {
	log.Printf("Max was invoked")
	maxNumber := int32(0)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}
		res := req.Number
		if res > maxNumber {
			maxNumber = res

			err = stream.Send(&proto.MaxResponse{
				Maximum: maxNumber,
			})
			if err != nil {
				log.Printf("Error while sending data to client: %v\n", err)
			}
		}
	}
}
