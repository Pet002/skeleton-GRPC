package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/promptlabth/orch_srver/app/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	calculator.RegisterCalculatorServer(grpcServer, calculator.NewCalculatorServer())

	reflection.Register(grpcServer)

	fmt.Println("start serve @" + port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
