package main

import (
	"log"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}
	defer conn.Close()
	c := proto.NewGreetServiceClient(conn)
	doGreetManyTimes(c)
}
