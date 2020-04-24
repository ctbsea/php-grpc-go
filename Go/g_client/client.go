package main

import (
	"context"
	"log"
	"os"
	pb "rpcx/pb"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	var age int64
	if len(os.Args) > 1 {
		name = os.Args[1]
		age ,_  = strconv.ParseInt(os.Args[2] , 10 , 32)
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name , Age: int32(age)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
	r2, err := c.Msg(ctx, &pb.MsgReq{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r2.GetMsg())
	r3, err := c.Msg(ctx, &pb.MsgReq{})
	log.Printf("Greeting: %s", r3.GetMsg())
}