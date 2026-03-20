package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type RestClient struct {
	BaseURL string
	Client  *http.Client
}

func NewRestClient(url string) *RestClient {
	return &RestClient{
		BaseURL: url,
		Client: &http.Client{
			Timeout: 5 * time.Second,
		},
	}
}

func (c *RestClient) SendRequest(payload interface{}) (time.Duration, error) {
	start := time.Now()

	data, _ := json.Marshal(payload)

	resp, err := c.Client.Post(c.BaseURL+"/process", "application/json", bytes.NewBuffer(data))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	return time.Since(start), nil
}