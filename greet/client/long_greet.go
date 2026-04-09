package main

import (
	"context"
	"log"
	"time"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
)

func doLongGreet(c proto.GreetServiceClient) {
	log.Println("doLongGreet was invoked")
	reqs := []*proto.GreetRequest{
		{FirstName: "Álvaro"},
		{FirstName: "Joina"},
		{FirstName: "Teste"},
	}
	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Error while calling LongGreetL %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while recieving response from LongGreet %v\n", err)
	}
	log.Printf("LongGreet: %s\n", res.Result)
}
