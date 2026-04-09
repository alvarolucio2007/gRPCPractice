package main

import (
	"context"
	"io"
	"log"
	"time"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
)

func doGreetEveryone(c proto.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")
	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v\n", err)
	}
	reqs := []*proto.GreetRequest{
		{FirstName: "Álvaro"},
		{FirstName: "Joina"},
		{FirstName: "Mainha"},
	}
	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error while recieving %v\n", err)
				break
			}
			log.Printf("Recieved %v\n", res.Result)
		}
		close(waitc)
	}()
	<-waitc
}
