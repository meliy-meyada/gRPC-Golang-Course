package main

import (
	"log"

	pb "github.com/meliy-meyada/gRPC-Golang-Course/calculator/proto"
)



func (s *Server) Primes(in *pb.PrimesRequest, stream pb.CalculatorService_PrimesServer) error{
	log.Printf("Primes Function was invoked with %v\n" , in)

	number := in.Number
	divisor := int64(2)

	for number > 1{
		if number % divisor == 0 {
			stream.Send(&pb.PrimesResponse{
				Result: divisor,
			})

			number /= divisor
		}else{
			divisor++
		}
	}
	return nil
}