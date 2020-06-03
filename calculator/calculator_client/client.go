package main

import (
	"context"
	"fmt"
	"groc-go/calculator/calculatorpb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Helo, I'm the client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Couldn't connect: %v", err)
	}

	defer cc.Close()

	c := calculatorpb.NewCalculatorServiceClient(cc)

	doUnary(c)
}

func doUnary(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Unary PRC...")
	req := &calculatorpb.CalculatorRequest{
		Numbers: &calculatorpb.Numbers{
			X: 10,
			Y: 3,
		},
	}

	res, err := c.Calculate(context.Background(), req)

	if err != nil {
		log.Fatalf("Error from server: %v", err)
	}

	log.Printf("Response from calculate: %v", res)
}
