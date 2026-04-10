package main

import (
	"context"
	"io"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/max/proto"
)

func doMax(c proto.MaxServiceClient) {
	log.Println("doMax was invoked")
	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream %v\n", err)
	}
	reqs := []*proto.MaxRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			if err := stream.Send(req); err != nil {
				log.Printf("Error %v\n", err)
			}
		}
		if err := stream.CloseSend(); err != nil {
			log.Printf("Error %v\n", err)
		}
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while receiving %v\n", err)
				break
			}
			log.Printf("Recieved %v\n", res.Maximum)
		}
		close(waitc)
	}()
	<-waitc
}
