package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/service"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router){
	router.HandleFunc("/api/getservices", GetServices).Methods("GET")
	router.HandleFunc("/api/saveservice", SaveService).Methods("POST")
	router.HandleFunc("/api/deleteservice/{id}", DeleteServiceByID).Methods("DELETE")
}
// GetServices Read services
func GetServices(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	log.Info("This is GetServices handler")
	serviceData, err := service.GetServices(ctx)
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
		Description: "Service records successfully",
		Data:        serviceData,
	}
	utils.Render(w, response)
	log.Infof("serviceData : %v ", serviceData)

}

func SaveService(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	log.Info("This is Service handler")
	// Insert consumption data
	var serviceRecord models.Services
	err := json.NewDecoder(r.Body).Decode(&serviceRecord)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	log.Infof("service good************: %v ", serviceRecord)
	serviceData, err := service.SaveService(ctx, serviceRecord)
	if err != nil {
		log.Info("you are in SaveService handler err block")
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	log.Info("you are in SaveService handler out of err block")
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusCreated,
		Description: "Service record is inserted successfully",
		Data:        serviceData,
	}
	utils.Render(w, response)
	log.Infof("ServiceList : success data %v ", serviceRecord)
}

func DeleteServiceByID(w http.ResponseWriter, r *http.Request) {
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
	log.Infof("pkID : %v ", pkID)
	rowsEffected, err := service.DeleteServiceByID(ctx, pkID)
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
		Description: fmt.Sprintf(" %d rows deleted from Services table", rowsEffected),
	}
	log.Infof("GetServices : error %v ", err)
	utils.Render(w, response)
}
