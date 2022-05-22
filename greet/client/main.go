package main

import (
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/meliy-meyada/gRPC-Golang-Course/greet/proto"
)


var addr string = "localhost:50051"

func main(){
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)

	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)


	// doGreed(c)
	// doGreetManyTimes(c)
	// doLongGreet(c)
	// doGreetEveryone(c)
	// doGreetWithDeadline(c, 1*time.Second)
	doGreetWithDeadline(c, 5*time.Second)
}