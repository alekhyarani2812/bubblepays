package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	consumptionApi "github.com/bubblepays/api/handler/consumption"
	"github.com/bubblepays/models"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func SkipTestGetConsumptions(t *testing.T){
	//req, err := http.NewRequest("GET", "/items", nil)
	// assert.NoError(t, err)
	// rr := httptest.NewRecorder()
	// handler := http.HandlerFunc(getItems)
	// handler.ServeHTTP(rr, req)
	// assert.Equal(t, http.StatusOK, rr.Code)
	// var responseItems []Item
	// err = json.NewDecoder(rr.Body).Decode(&responseItems)
	// assert.NoError(t, err)
	// assert.Equal(t, len(items), len(responseItems))
	req, err := http.NewRequest("GET", "/api/consumptions", nil)
	assert.NoError(t, err)
    rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/consumptions/{id}",consumptionApi.GetConsumptions)
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var responseConsumptions []models.Consumption
	err = json.NewDecoder(rr.Body).Decode(&responseConsumptions)
	assert.NoError(t, err)
	//assert.Equal(t, len(), len(responseConsumptions))
}
func SkipTestGetConsumptionByID(t *testing.T){
	req, err := http.NewRequest("GET", "/api/consumptions/{id}", nil)
	assert.NoError(t, err)
    rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/api/consumptions/{id}",consumptionApi.GetConsumptions)
	router.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)
	var responseConsumption []models.Consumption
	err = json.NewDecoder(rr.Body).Decode(&responseConsumption)
	assert.NoError(t, err)
	assert.Equal(t, 7, len(responseConsumption))
}