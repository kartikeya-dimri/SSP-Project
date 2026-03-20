package main

import (
	"log"
	"net/http"

	"rest-vs-grpc-benchmark/rest/router"
)

func main() {
	router.SetupRouter()

	port := ":8080"
	log.Println("REST server running on port", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal(err)
	}
}