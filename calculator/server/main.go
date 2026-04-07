package calculator

import (
	"log"
	"net"
)

var addr string = "0.0.0.0:50051"

type Server struct{}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Failed to listen on %v", err)
	}
}
