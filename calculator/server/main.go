package calculator

import (
	"log"
	"net"

	"github.com/alvarolucio2007/gRPCPractice/calculator/proto"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	proto.SumResponse
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v", err)
	}
}
