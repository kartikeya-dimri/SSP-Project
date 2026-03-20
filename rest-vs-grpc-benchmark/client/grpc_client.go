package main

import (
	"context"
	"time"

	"google.golang.org/grpc"

	"rest-vs-grpc-benchmark/grpc/pb"
)

type GrpcClient struct {
	Client pb.BenchmarkServiceClient
}

func NewGrpcClient(addr string) (*GrpcClient, error) {
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewBenchmarkServiceClient(conn)

	return &GrpcClient{Client: client}, nil
}

func (c *GrpcClient) SendRequest(payload *pb.RequestPayload) (time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	start := time.Now()

	_, err := c.Client.Process(ctx, payload)
	if err != nil {
		return 0, err
	}

	return time.Since(start), nil
}