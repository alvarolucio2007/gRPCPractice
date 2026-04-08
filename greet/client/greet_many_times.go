package main

import (
	"context"
	"io"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
)

func doGreetManyTimes(c proto.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")
	req := &proto.GreetRequest{
		FirstName: "Álvaro",
	}
	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes:%v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error while reading the stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
		doGreetManyTimes(c)
	}
}
