package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"

	"rest-vs-grpc-benchmark/grpc/pb"
)

const (
	NUM_REQUESTS = 1000
	CONCURRENCY  = 50
)

func generatePayload() *pb.RequestPayload {
	return &pb.RequestPayload{
		UserId:  1,
		Username: "test",
		Data: []*pb.NestedData{
			{Id: 1, Name: "item1", Active: true, Tags: []string{"a", "b"}},
			{Id: 2, Name: "item2", Active: false, Tags: []string{"c", "d"}},
		},
	}
}

func calculatePercentile(latencies []float64, p float64) float64 {
	sort.Float64s(latencies)
	index := int(math.Ceil(p/100.0*float64(len(latencies)))) - 1
	return latencies[index]
}

func saveToCSV(filename string, latencies []float64) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	writer.Write([]string{"latency_ms"})

	for _, l := range latencies {
		writer.Write([]string{strconv.FormatFloat(l, 'f', 2, 64)})
	}
}

func runREST() {
	fmt.Println("Running REST benchmark...")

	rc := NewRestClient("http://localhost:8080")

	var latencies []float64
	var mu sync.Mutex
	var wg sync.WaitGroup

	startTotal := time.Now()

	for i := 0; i < NUM_REQUESTS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			payload := generatePayload()

			latency, err := rc.SendRequest(payload)
			if err == nil {
				mu.Lock()
				latencies = append(latencies, float64(latency.Milliseconds()))
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	totalTime := time.Since(startTotal).Seconds()

	printStats(latencies, totalTime)

	saveToCSV("../results/rest_latencies.csv", latencies)
}

func runGRPC() {
	fmt.Println("Running gRPC benchmark...")

	gc, err := NewGrpcClient("localhost:50051")
	if err != nil {
		panic(err)
	}

	var latencies []float64
	var mu sync.Mutex
	var wg sync.WaitGroup

	startTotal := time.Now()

	for i := 0; i < NUM_REQUESTS; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			payload := generatePayload()

			latency, err := gc.SendRequest(payload)
			if err == nil {
				mu.Lock()
				latencies = append(latencies, float64(latency.Milliseconds()))
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	totalTime := time.Since(startTotal).Seconds()

	printStats(latencies, totalTime)

	saveToCSV("../results/grpc_latencies.csv", latencies)
}

func printStats(latencies []float64, totalTime float64) {
	p50 := calculatePercentile(latencies, 50)
	p95 := calculatePercentile(latencies, 95)
	p99 := calculatePercentile(latencies, 99)

	throughput := float64(len(latencies)) / totalTime

	fmt.Println("Results:")
	fmt.Printf("p50 latency: %.2f ms\n", p50)
	fmt.Printf("p95 latency: %.2f ms\n", p95)
	fmt.Printf("p99 latency: %.2f ms\n", p99)
	fmt.Printf("Throughput: %.2f req/sec\n", throughput)
	fmt.Println("----------------------------------")
}

func main() {
	runREST()
	runGRPC()
}