package main

import (
	"context"
	"fmt"
	"groc-go/greet/greetpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello, I'm client")
	// Open an insecure connection to localhost:50051
	// grpc.WithInsecure() to tell grpc to not to use ssl (for now)
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	// defer is used to execute at the very end (Close the connection)
	defer cc.Close()

	c := greetpb.NewGreetServiceClient(cc)

	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Unary RPC...")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Sathish",
			LastName:  "S",
		},
	}
	res, err := c.Greet(context.Background(), req)

	if err != nil {
		log.Fatalf("Error from server: %v", err)
	}

	log.Printf("Response from greet: %v", res)
}
