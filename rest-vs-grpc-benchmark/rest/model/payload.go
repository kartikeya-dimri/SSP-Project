package model

type NestedData struct {
	ID     int      `json:"id"`
	Name   string   `json:"name"`
	Active bool     `json:"active"`
	Tags   []string `json:"tags"`
}

type RequestPayload struct {
	UserID   int          `json:"user_id"`
	Username string       `json:"username"`
	Data     []NestedData `json:"data"`
}

type ResponsePayload struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}