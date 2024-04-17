package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bubblepays/models"
	serviceenrollment "github.com/bubblepays/repository/service_enrollment"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

func GetServiceEnrollment(w http.ResponseWriter, r *http.Request) {
	ctx := utils.WithTransactionID(r.Context())
	serviceEnrollRecords, err := serviceenrollment.GetServiceEnrollment(ctx)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    serviceEnrollRecords,
	}
	utils.Render(w, response)
}

func GetServiceEnrollmentById(w http.ResponseWriter, r *http.Request) {
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
	serviceEnrollmentRecord, err := serviceenrollment.GetServiceEnrollmentById(ctx, pkId)
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
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    serviceEnrollmentRecord,
	}
	utils.Render(w, response)
}

func SaveServiceEnrollment(w http.ResponseWriter, r *http.Request) {
	ctx := utils.WithTransactionID(r.Context())
	var serviceEnrollment models.ServiceEnrollment
	err := json.NewDecoder(r.Body).Decode(&serviceEnrollment)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	serviceEnrollmentRecord, err := serviceenrollment.SaveServiceEnrollment(ctx, serviceEnrollment)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: fmt.Sprintf("SE record with id: %d inserted successfully", serviceEnrollmentRecord.Id),
		Data:        serviceEnrollmentRecord,
	}
	utils.Render(w, response)
}

func UpdateServiceEnrollment(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("pkId : ", pkId)
	var serviceEnrollment models.ServiceEnrollment
	if err := json.NewDecoder(r.Body).Decode(&serviceEnrollment); err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	serviceEnrollmentRecord, err := serviceenrollment.UpdateServiceEnrollment(ctx, pkId, serviceEnrollment)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    serviceEnrollmentRecord,
	}
	utils.Render(w, response)
}

func DeleteServiceEnrollment(w http.ResponseWriter, r *http.Request) {
	ctx := utils.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	idStr := vars["id"]
	pkId, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	fmt.Println("pkId : ", pkId)
	rowsEffected, err := serviceenrollment.DeleteServiceEnrollment(ctx, pkId)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	if pkId == 0 {
		response := models.Response{
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf("you should not enter id %d : ", pkId),
		}
		utils.Render(w, response)
		return
	}
	if rowsEffected == 0 {
		response := models.Response{
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf(" please enter sufficient id %d : ", pkId),
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: fmt.Sprintf("record with id %d has deleted successfully ", pkId),
	}
	utils.Render(w, response)
}
