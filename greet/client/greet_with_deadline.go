package main

import (
	"context"
	"log"
	"time"

	"github.com/alvarolucio2007/gRPCPractice/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c proto.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	req := &proto.GreetRequest{
		FirstName: "Álvaro",
	}
	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded")
				return
			}
			log.Fatalf("Unexpected gRPC error: %v\n", err)

		} else {
			log.Fatalf("a non gRPC error: %v\n", err)
		}
	}
	log.Printf("GreetWithDeadline: %s\n", res.Result)
}
