package main

import (
	"context"
	"fmt"
	"groc-go/calculator/calculatorpb"
	"io"
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
	doServerStreaming(c)
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

func doServerStreaming(c calculatorpb.CalculatorServiceClient) {
	fmt.Println("Starting to do Server Streaming RPC...")

	req := &calculatorpb.PrimeNumberDecompositionRequest{
		Number: &calculatorpb.Number{
			X: 210,
		},
	}

	resStream, err := c.PrimeNumberDecomposition(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}

	for {
		msg, err := resStream.Recv()
		if err == io.EOF {
			// We reached the end of the stream
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Response: %v", msg.GetResult())
	}
}
