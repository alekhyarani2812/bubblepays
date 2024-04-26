package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/service"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

// Read services
func GetServices(w http.ResponseWriter, r *http.Request) {
	ctx := utils.WithTransactionID(r.Context())
	fmt.Println("This is GetServices handler")
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
	fmt.Println("serviceData : ", serviceData)

}

func SaveService(w http.ResponseWriter, r *http.Request) {
	ctx := utils.WithTransactionID(r.Context())
	fmt.Println("This is Service handler")
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
	log.Println("service good************: ", serviceRecord)
	serviceData, err := service.SaveService(ctx, serviceRecord)
	if err != nil {
		log.Println("you are in SaveService handler err block")
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	log.Println("you are in SaveService handler out of err block")
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusCreated,
		Description: "Service record is inserted successfully",
		Data:        serviceData,
	}
	utils.Render(w, response)
	log.Println("ServiceList : success data ", serviceRecord)
}

func DeleteServiceById(w http.ResponseWriter, r *http.Request) {
	ctx := utils.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	idStr := vars["id"]
	pkId, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	log.Println("pkId : ", pkId)
	rowsEffected, err := service.DeleteServiceById(ctx, pkId)
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
	log.Printf("GetServices : error %v ", err)
	utils.Render(w, response)
}
