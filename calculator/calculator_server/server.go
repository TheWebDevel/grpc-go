package main

import (
	"context"
	"fmt"
	"groc-go/calculator/calculatorpb"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct{}

// Calculate function
func (*server) Calculate(ctx context.Context, req *calculatorpb.CalculatorRequest) (*calculatorpb.CalculatorResponse, error) {
	fmt.Printf("Calculate function was invoked with: %v\n", req)
	x := req.GetNumbers().GetX()
	y := req.GetNumbers().GetY()
	result := x + y

	res := &calculatorpb.CalculatorResponse{
		Result: result,
	}

	return res, nil
}

// Prime number decomposition
func (*server) PrimeNumberDecomposition(req *calculatorpb.PrimeNumberDecompositionRequest, stream calculatorpb.CalculatorService_PrimeNumberDecompositionServer) error {
	fmt.Printf("PrimeNumberDecomposition was invoked with %v", req)

	n := req.GetNumber().GetX()

	var k int64 = 2
	for n > 1 {
		if n%k == 0 {
			result := k
			res := &calculatorpb.PrimeNumberDecompositionResponse{
				Result: result,
			}

			stream.Send(res)
			time.Sleep(1000 * time.Millisecond)
			n = n / k
		} else {
			k = k + 1
		}
	}
	return nil
}

func main() {
	fmt.Println("Hello world")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
