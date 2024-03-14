package calculator

import context "context"

type calculatorServer struct {
	UnimplementedCalculatorServer
}

func NewCalculatorServer() CalculatorServer {
	return &calculatorServer{}
}

func (c *calculatorServer) Hello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {

	return &HelloResponse{
		Message: "hello " + req.Name + " " + req.Surname,
	}, nil
}
