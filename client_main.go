package main

import (
	"google.golang.org/grpc"
	"log"
	pb "github.com/grpcpractice/service"
	"context"
	"fmt"
)

const (
	address  = "localhost:50001"
	defaultName = "World"
)

func AddUser() *pb.Person{
	x := pb.Person{
		FirstName:    "John",
		LastName:     "Doe",
		DateOfBirth:  "1960-10-17T0:00:00Z",
		Cool:         true,
		ArgumentsWon: 7,
		Hobby : []*pb.Hobby{
			{
				Name:        "Running",
				Description: "Occasionally, about 10km a week",
			}, {
				Name:        "Computer games",
				Description: "Flappy bird, mostly",
			},
		},
	}

	return &x
}

func main()  {
	conn , err := grpc.Dial(address,grpc.WithInsecure())

	if err!=nil{
		log.Fatalf("Did not connect : %v",err)
	}

	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	person := AddUser()
	name := person.FullName()

	r , err := c.SayHello(context.Background(),&pb.HelloRequest{Name:name})

	if err!=nil{
		log.Fatalf("Could not greet %v",err)
	}

	log.Printf("Greeting %s",r.Message)
	fmt.Println(person)
}