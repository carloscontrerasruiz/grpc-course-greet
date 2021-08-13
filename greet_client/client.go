package main

import (
	"context"
	"fmt"
	"log"

	gp "github.com/carloscontrerasruiz/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello Im the client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not conect: %v", err)
	}

	defer cc.Close()

	c := gp.NewGreetServiceClient(cc)

	doUnary(c)

}

func doUnary(c gp.GreetServiceClient) {
	req := &gp.GreetRequest{
		Greeting: &gp.Greeting{
			FirstName: "Carlos111",
			LastName:  "CC",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling greet %v", err)
	}

	log.Printf("Response from greeting %v", res.Result)
}
