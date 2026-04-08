package main

import (
	"context"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/calculator/proto"
)

func doCalculate(c proto.SumServiceClient) {
	log.Println("doCalculate was invoked")
	res, err := c.Sum(context.Background(), &proto.SumRequest{
		FirstNumber: 25, SecondNumber: 999999999999999999,
	})
	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}
	log.Printf("Greeting: %v\n", res.Result)
}
