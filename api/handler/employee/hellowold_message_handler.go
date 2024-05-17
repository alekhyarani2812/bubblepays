package api

import (
	"encoding/json"
	bcontext "github.com/bubblepays/context"
	"log"
	"net/http"
)

type message struct {
	Text string `json:"name"`
}

func Hello(w http.ResponseWriter, r *http.Request) {
	// Get the context from the request
	ctx := r.Context()
	w.Write([]byte("Hello transactionID: " + bcontext.GetTransactionID(ctx)))
}

func GetMessageHandler(w http.ResponseWriter, r *http.Request) {
	message := message{Text: "***********Hello, World!**************"}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(message)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
