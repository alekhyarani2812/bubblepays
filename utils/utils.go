package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
	"todoproject/models"
)

func Render(w http.ResponseWriter, response models.Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)
	//json.NewEncoder(w).Encode(response)
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatal(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// func string to int
// employeeIDStr :=  r.PathValue("id")
func StringToInt(idStr string) (int, error) {
	if idStr == "" {
		return 0, errors.New("id should not be empty")
	}
	// Convert employeeIDStr to int
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println("getting error while strconv err:", err)
		//return 0, errors.New(fmt.Sprintf("given id (%v) is not a number", idStr))

		return 0, fmt.Errorf("given id (%v) is not a number", idStr)
	}
	return id, nil
}
func StringToTime(str string) (*time.Time,error){
	// Parse the string into a time.Time object
	// Define the layout of the string representing the time
	layout := "2006-01-02 15:04:05"
	parsedTime, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing time:", err)
		return nil,err
	}
	// Print the parsed time
	fmt.Println("Parsed time:", parsedTime)
	return &parsedTime,nil
}
