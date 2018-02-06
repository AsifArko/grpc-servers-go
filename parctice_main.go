package main

import (
	"github.com/grpcpractice/practice"
	"fmt"
)

func main() {
	p := protofile.Person{
		FirstName:    "John",
		LastName:     "Doe",
		DateOfBirth:  "1960-10-17T0:00:00Z",
		Cool:         true,
		ArgumentsWon: 7,
		Hobby : []*protofile.Hobby{
			{
				Name:        "Running",
				Description: "Occasionally, about 10km a week",
			}, {
				Name:        "Computer games",
				Description: "Flappy bird, mostly",
			},
		},
	}

	fmt.Printf("Person created for .proto structure: %v\n", p)
	fmt.Println(p.Fullname())
}

