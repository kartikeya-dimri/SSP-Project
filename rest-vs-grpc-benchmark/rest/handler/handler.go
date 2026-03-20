package handler

import (
	"encoding/json"
	"net/http"

	"rest-vs-grpc-benchmark/rest/model"
)

func ProcessHandler(w http.ResponseWriter, r *http.Request) {
	var req model.RequestPayload

	// Decode JSON request
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// Simulate processing (important for realism)
	for _, d := range req.Data {
		_ = d.ID
	}

	resp := model.ResponsePayload{
		Status:  "success",
		Message: "Processed successfully",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}