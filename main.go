package main

import (
	"fmt"
	pb "grpc-demo/proto"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type HelloService struct {
	pb.UnimplementedHelloServiceServer
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()

	reflection.Register(s)
	pb.RegisterHelloServiceServer(s, &HelloService{})

	fmt.Println("server is running at port 8080")

	if err := s.Serve(l); err != nil {
		panic(err)
	}
}
