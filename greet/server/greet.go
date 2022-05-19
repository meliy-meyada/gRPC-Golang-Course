package main

import (
	"context"
	"log"

	pb "github.com/meliy-meyada/gRPC-Golang-Course/greet/proto"
)



func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet Function wa invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello" + in.FirstName,
	}, nil
}