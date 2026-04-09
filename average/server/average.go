package main

import (
	"io"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/average/proto"
)

func (s *Server) Average(stream proto.AverageService_AverageServer) error {
	log.Println("Average function was invoked")
	res, counter := 0, 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.AverageResponse{
				Number: float32(res / counter),
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}
		log.Printf("Receiving %v\n", req)

		res += int(req.Number)
		counter += 1
	}
}
