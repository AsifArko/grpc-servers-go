package main

import (
	pb "github.com/grpcpractice/userprofile"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

const(
	address = "localhost:12345"
)

func AddProfile() *pb.ProfReq{
	profile := &pb.ProfReq{Fname:"Asif",Lname:"Arko",Dob:"3/3/1993",Hobby:[]*pb.Hobbies{
		{
		Name:"gaming",
		Description:"occationally gaming",
		},
		{
			Name:"Running",
			Description:"occationally Running",
		},
	}}
	return profile
}

func main()  {
	conn , err := grpc.Dial(address,grpc.WithInsecure())

	if err != nil{
		fmt.Println(err.Error(),"Did not connect .")
	}

	defer conn.Close()

	c := pb.NewCreateUserClient(conn)

	profile := AddProfile()

	r , err := c.HandleProfReq(context.Background(),profile)

	if err != nil{
		log.Fatalf("Could not handle the profile request")
	}

	fmt.Println(r)
}