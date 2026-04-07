package main

import (
	"context"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
)

func doGreet(c proto.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &proto.GreetRequest{
		FirstName: "Álvaro",
	})
	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}
	log.Printf("Greeting: %s\n", res.Result)
}
