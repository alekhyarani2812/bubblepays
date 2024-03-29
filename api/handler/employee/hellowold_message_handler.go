package api

import (
	"encoding/json"
	"log"
	"net/http"
)

type Message struct {
	Text string `json:"name"`
}

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{Text: "Hello, World!"}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
