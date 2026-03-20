package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"rest-vs-grpc-benchmark/grpc/pb"
	"rest-vs-grpc-benchmark/grpc/server"
)

func main() {
	port := ":50051"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterBenchmarkServiceServer(grpcServer, &server.Server{})

	log.Println("gRPC server running on port", port)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}