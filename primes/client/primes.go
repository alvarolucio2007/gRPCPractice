package main

import (
	"context"
	"io"
	"log"

	"github.com/alvarolucio2007/gRPCPractice/primes/proto"
)

func doPrimesManyTimes(c proto.PrimeServiceClient) {
	log.Println("doPrimes was invoked")
	req := &proto.Primes{
		Number: 210,
	}
	stream, err := c.PrimeManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling PrimeManyTimes: %v\n", err)
	}
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("Error while reading the stream: %v\n ", err)
		}
		log.Printf("PrimeManyTimes: %s\n", msg.ResultFactored)
	}
}
