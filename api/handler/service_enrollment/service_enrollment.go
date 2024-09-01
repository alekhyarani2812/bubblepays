package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	serviceenrollment "github.com/bubblepays/repository/service_enrollment"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)
// LoadRoutes for routing all service enrollment api's
func LoadRoutes(router *mux.Router){
	router.HandleFunc("/api/getserviceenrollment", GetServiceEnrollment).Methods("GET")
	router.HandleFunc("/api/getserviceenrollmentbyid/{id}", GetServiceEnrollmentByID).Methods("GET")
	router.HandleFunc("/api/saveserviceenrollment", SaveServiceEnrollment).Methods("POST")
	router.HandleFunc("/api/updateserviceenrollment/{id}", UpdateServiceEnrollment).Methods("PUT")
	router.HandleFunc("/api/deleteserviceenrollment/{id}", DeleteServiceEnrollment).Methods("DELETE")
}
// GetServiceEnrollment retrieves all serviceEnrollment records from the table
func GetServiceEnrollment(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
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
// GetServiceEnrollmentByID Retrieve the serviceEnrollments record based on the serviceEnrollmentID
func GetServiceEnrollmentByID(w http.ResponseWriter, r *http.Request) {
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
	serviceEnrollmentRecord, err := serviceenrollment.GetServiceEnrollmentByID(ctx, pkID)
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
// SaveServiceEnrollment for inserting serviceEnrollments into table
func SaveServiceEnrollment(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
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
		Description: fmt.Sprintf("SE record with id: %d inserted successfully", serviceEnrollmentRecord.ID),
		Data:        serviceEnrollmentRecord,
	}
	utils.Render(w, response)
}
// UpdateServiceEnrollment function is used to update the data in the serviceEnrollment table
func UpdateServiceEnrollment(w http.ResponseWriter, r *http.Request) {
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
	serviceEnrollmentRecord, err := serviceenrollment.UpdateServiceEnrollment(ctx, pkID, serviceEnrollment)
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
// DeleteServiceEnrollment for deleting serviceenrollment record based on serviceenrollmentId
func DeleteServiceEnrollment(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	idStr := vars["id"]
	pkID, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	log.Infof("pkID : %v ", pkID)
	rowsEffected, err := serviceenrollment.DeleteServiceEnrollment(ctx, pkID)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	if pkID == 0 {
		response := models.Response{
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf("you should not enter id %d : ", pkID),
		}
		utils.Render(w, response)
		return
	}
	if rowsEffected == 0 {
		response := models.Response{
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf(" please enter sufficient id %d : ", pkID),
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: fmt.Sprintf("record with id %d has deleted successfully ", pkID),
	}
	utils.Render(w, response)
}
