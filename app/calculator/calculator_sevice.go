package calculator

import (
	context "context"
	"log"
)

type calculatorServer struct {
	UnimplementedCalculatorServer
}

func NewCalculatorServer() CalculatorServer {
	return &calculatorServer{}
}

func (c *calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	log.Println("logging: " + req.Name)
	return &HelloResponse{
		Message: "hello " + req.Name + " " + req.Surname,
	}, nil
}
