package models

// Response model for utils
type Response struct {
	Message     string      `json:"message"`
	Description string      `json:"description"`
	Error       string      `json:"error"`
	Status      int         `json:"status"`
	Data        interface{} `json:"data"`
}
