package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/consumption"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router){
	router.HandleFunc("/api/consumptions", GetConsumptions).Methods("GET")
	router.HandleFunc("/api/consumptions/{id}", GetConsumptionByID).Methods("GET")
	router.HandleFunc("/api/saveconsumption", SavesConsumption).Methods("POST")
	router.HandleFunc("/api/deleteconsumption/{id}", DeleteConsumption).Methods("DELETE")
}
func GetConsumptions(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	consumptionList, err := consumption.GetConsumptions(ctx)
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
		Status:  http.StatusOK,
		Data:    consumptionList,
	}
	utils.Render(w, response)
}

//GetConsumptionByID Get 1 Consumption by ID
func GetConsumptionByID(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	idStr := vars["id"]
	//idStr := r.PathValue("id")
	pkID, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	consumptionList, err := consumption.GetConsumptionByID(ctx, pkID)
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
		Status:  http.StatusOK,
		Data:    consumptionList,
	}
	utils.Render(w, response)
}

//SavesConsumption Insert consumption data
func SavesConsumption(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	var consum models.Consumption
	err := json.NewDecoder(r.Body).Decode(&consum)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		log.Infot(ctx, "error StatusBadRequest in json.NewDecoder(r.Body).Decode(&consum)")
		return
	}
	consumData := consum
	consumptionRecord, err := consumption.SavesConsumption(ctx, consumData)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		log.Infot(ctx, "error StatusBadRequest in consumptionRecord")
		return
	}
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: "consumption record is inserted successfully",
		Data:        consumptionRecord,
	}
	utils.Render(w, response)
}

func DeleteConsumption(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	idStr := vars["id"]
	//idStr := r.PathValue("id")
	pkID, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	rowsEffected, err := consumption.DeleteConsumption(ctx, pkID)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	if rowsEffected != 0 {
		response := models.Response{
			Message:     utils.Success,
			Status:      http.StatusOK,
			Description: fmt.Sprintf("consumption record %v is deleted", pkID),
		}
		utils.Render(w, response)
	} else {
		response := models.Response{
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf("consumption record %v does not exist in consumption table", pkID),
		}
		utils.Render(w, response)
	}
}
