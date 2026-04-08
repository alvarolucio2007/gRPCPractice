package main

import (
	"log"

	"github.com/alvarolucio2007/gRPCPractice/primes/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v\n", err)
	}
	defer conn.Close()
	c := proto.NewPrimeServiceClient(conn)
	doPrimesManyTimes(c)
}
