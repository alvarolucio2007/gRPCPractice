package main

import (
	"log"

	"github.com/alvarolucio2007/gRPCPractice/primes/proto"
)

func (s *Server) PrimeManyTimes(in *proto.Primes, stream proto.PrimeService_PrimeManyTimesServer) error {
	log.Printf("Fatoração iniciada para: %d\n", in.Number)

	var starting uint32 = 2
	original := in.Number

	for original > 1 {
		if original%starting == 0 {
			// Se é divisível, manda o fator e divide o número
			if err := stream.Send(&proto.PrimeResponse{
				ResultFactored: starting,
			}); err != nil {
				return err
			}
			original = original / starting
		} else {
			// SE NÃO FOR DIVISÍVEL, incrementa o DIVISOR, não o número!
			starting++
		}
	}
	return nil
}
