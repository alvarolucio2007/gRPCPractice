package main

import (
	"context"
	"log"
	"time"

	"github.com/alvarolucio2007/gRPCPractice/average/proto"
)

func doAverage(c proto.AverageServiceClient) {
	log.Println("doAverage was invoked")
	reqs := []*proto.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
	}
	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("Error while calling doAverage %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		err := stream.Send(req)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
		time.Sleep(1 * time.Second)

	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while receiving response from Average %v\n", err)
	}
	log.Printf("Average: %f\n", res.Number)
}
