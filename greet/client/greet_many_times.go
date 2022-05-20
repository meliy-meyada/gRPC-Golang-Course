package main

import (
	"context"
	"io"
	"log"

	pb "github.com/meliy-meyada/gRPC-Golang-Course/greet/proto"
)


func doGreetManyTimes(c pb.GreetServiceClient){
	log.Println("doGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Meyada",
	}
	stream , err := c.GreetManyTimes(context.Background(), req)

	if err != nil{
		log.Fatalf("Error while calling GreetManyTimes: %v\n", err)
	}
	for {
		msg, err := stream.Recv()

		if err == io.EOF{
			break
		}
		if err != nil {
			log.Fatal("Error while Reading the stream: %v\n", err)
		}
		log.Printf("GreetManyTimes: %s\n", msg.Result)
	}
}