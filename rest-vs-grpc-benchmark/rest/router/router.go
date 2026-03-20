package router

import (
	"net/http"

	"rest-vs-grpc-benchmark/rest/handler"
)

func SetupRouter() {
	http.HandleFunc("/process", handler.ProcessHandler)
}