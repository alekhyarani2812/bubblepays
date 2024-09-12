package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"time"

	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
)

// Render for reading and giving json response
func Render(w http.ResponseWriter, response models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	// ctx := r.Context
	// w.Write([]byte(GetTransactionID(ctx)))
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// StringToInt IDStr := r.PathValue("id")
func StringToInt(idStr string) (int, error) {
	if idStr == "" {
		return 0, errors.New("id should not be empty")
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Infof("getting error while strconv err: %v ", err)
		return 0, fmt.Errorf("given id (%v) is not a number", idStr)
	}
	return id, nil
}
// StringToTime to convert string to time
func StringToTime(str string) (*time.Time, error) {
	// Parse the string into a time.Time object
	// Define the layout of the string representing the time
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		log.Errorf("Error parsing time: %v", err)
		return nil, err
	}
	log.Infof("Parsed time: %v", parsedTime)
	return &parsedTime, nil
}

func IsAlphanumeric(s string) bool {
	// Regular expression to check for alphanumeric characters
	regex := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	return regex.MatchString(s)
}

func IsValidCSVFilename(filename string) bool {
	// Regular expression to match alphanumeric characters, underscores, and end with .csv
	// regex := regexp.MustCompile(`^[a-zA-Z0-9_]+\.csv$`)
	// m100_11092024_183054
	regex := regexp.MustCompile(`^[a-zA-Z0-9]+_[0-9]{8}_[0-9]{6}\.csv$`)
	return regex.MatchString(filename)
}