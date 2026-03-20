package server

import (
	"context"

	"rest-vs-grpc-benchmark/grpc/pb"
)

type Server struct {
	pb.UnimplementedBenchmarkServiceServer
}

func (s *Server) Process(ctx context.Context, req *pb.RequestPayload) (*pb.ResponsePayload, error) {

	// Simulate processing (same as REST)
	for _, d := range req.Data {
		_ = d.Id
	}

	return &pb.ResponsePayload{
		Status:  "success",
		Message: "Processed successfully",
	}, nil
}