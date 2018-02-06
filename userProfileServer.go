package main

import (
	pb "github.com/grpcpractice/userprofile"
	"context"
	"fmt"
	"net"
	"os"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
)

const port = ":12345"

type server struct{}

func (s *server) HandleProfReq(ctx context.Context , inp *pb.ProfReq)(*pb.UserProf,error) {
		return &pb.UserProf{Prof:inp,Rc:&pb.RestraCoin{Coin:5}},nil;
}

func main()  {

	fmt.Println("Server is listening on localhost",port)

	lis , err := net.Listen("tcp",port)

	if err!=nil{
		fmt.Println(err.Error())
		os.Exit(1)
	}

	srv := grpc.NewServer()
	pb.RegisterCreateUserServer(srv,&server{})
	reflection.Register(srv)

	if err := srv.Serve(lis); err!=nil{
		log.Fatalf("Failed to serve %v",err)
	}

}

