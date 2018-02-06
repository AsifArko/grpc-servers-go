package main

import (
	pb "github.com/grpcpractice/service"
	"context"
	"net"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"

)

const port = ":50001"

type server struct {}

func (s *server) SayHello(ctx context.Context , in *pb.HelloRequest)(*pb.HelloReply,error){
	return &pb.HelloReply{Message:"hello "+in.Name},nil
}

func main()  {
	fmt.Println("Server is running on localhost",port)
	lis , err := net.Listen("tcp",port)

	if err != nil{
		fmt.Println(err.Error())
	}

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s,&server{})
	reflection.Register(s)

	if err := s.Serve(lis);err!=nil{
		log.Fatalf("Failed to serve %v",err)
	}

}