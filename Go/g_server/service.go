package main

//protoc --proto_path pb/ --go_out=plugins=grpc:pb helloword.proto
//proto_path  指定对导入文件的搜索路径
//go_out/paython_out/java_out 生成对应语言的代码

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	pb "rpcx/pb"
	"strconv"

	"google.golang.org/grpc"

)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v , %v", in.GetName() , in.GetAge())
	age := int64(in.GetAge())
	return &pb.HelloReply{Message: "Hello " + in.GetName() + ", age:" + strconv.FormatInt(age ,10)}, nil
}

// SayHello implements helloworld.GreeterServer
func (s *server) Msg(ctx context.Context, in *pb.MsgReq) (*pb.MsgRep, error) {
	num := int64(rand.Int())
	fmt.Print(1)
	return &pb.MsgRep{Msg: strconv.FormatInt(num ,10)}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}