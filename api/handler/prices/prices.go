package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/prices"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)
// LoadRoutes for routing all price api's
func LoadRoutes(router *mux.Router){
	router.HandleFunc("/api/getprices", GetPrices).Methods("GET")
	router.HandleFunc("/api/getpricesbyid/{id}", GetPricesByID).Methods("GET")
	router.HandleFunc("/api/saveprice", SavePrice).Methods("POST")
	router.HandleFunc("/api/updateprice/{id}", UpdatePrice).Methods("PUT")
	router.HandleFunc("/api/deleteprice/{id}", DeletePrice).Methods("DELETE")
}
// GetPrices retrieves all price records from the table
func GetPrices(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	priceReords, err := prices.GetPrices(ctx)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Status:  http.StatusBadRequest,
			Message: utils.Fail,
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Status:  http.StatusFound,
		Message: utils.Success,
		Data:    priceReords,
	}
	utils.Render(w, response)
}
// GetPricesByID Retrieve the prices record based on the priceId
func GetPricesByID(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	idStr := vars["id"]
	pkID, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	priceReord, err := prices.GetPricesByID(ctx, pkID)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusFound,
		Data:    priceReord,
	}
	utils.Render(w, response)
}
// SavePrice for inserting prices into table
func SavePrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	var price models.Price
	err := json.NewDecoder(r.Body).Decode(&price)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
		}
		utils.Render(w, response)
		return
	}
	priceRecord, err := prices.SavePrice(ctx, price)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusCreated,
		Description: fmt.Sprintf("%d record inserted successfully", priceRecord.ID),
		Data:        priceRecord,
	}
	utils.Render(w, response)
}
// UpdatePrice function is used to update the data in the price table
func UpdatePrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	idStr := vars["id"]
	pkID, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	log.Infoft(ctx, "pkID : %v ", pkID)
	var price models.Price
	err1 := json.NewDecoder(r.Body).Decode(&price)
	if err1 != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err1.Error(),
		}
		utils.Render(w, response)
		return
	}
	updatedRecord, err := prices.UpdatePrice(ctx, pkID, &price)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: fmt.Sprintf("%d record updated successfully", pkID),
		Data:        updatedRecord,
	}
	utils.Render(w, response)
}
// DeletePrice for deleting price record based on priceId
func DeletePrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	idStr := vars["id"]
	pkID, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	log.Infoft(ctx, "pkID : %v ", pkID)
	rowsEffected, err := prices.DeletePrice(ctx, pkID)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: fmt.Sprintf(" %d rows deleted from price table", rowsEffected),
	}
	utils.Render(w, response)
}
