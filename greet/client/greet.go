package main

import (
	"context"
	"log"

	pb "github.com/meliy-meyada/gRPC-Golang-Course/greet/proto"
)

func doGreed(c pb.GreetServiceClient){
	log.Println("doGreet was invoked")
	res , err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Meyada",
	})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Panicf("Greeting : %s\n", res.Result)
}