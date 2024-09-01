package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/plans"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)
//LoadRoutes for routing all plan api's
func LoadRoutes(router *mux.Router){
	router.HandleFunc("/api/getplans", GetPlans).Methods("GET")
	router.HandleFunc("/api/getplanbyid/{id}", GetPlanByID).Methods("GET")
	router.HandleFunc("/api/saveplan", SavePlan).Methods("POST")
	router.HandleFunc("/api/updateplan/{id}", UpdatePlan).Methods("PUT")
	router.HandleFunc("/api/deleteplan/{id}", DeletePlan).Methods("DELETE")
}
// GetPlans retrieves all plan records from the table
func GetPlans(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	plansList, err := plans.GetPlans(ctx)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		log.Infoft(ctx, "GetPlans : error %v ", err)
		return
	}
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    plansList,
	}
	utils.Render(w, response)
}
// GetPlanByID Retrieve the plans record based on the planId
func GetPlanByID(w http.ResponseWriter, r *http.Request) {
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
	planRecord, err := plans.GetPlanByID(ctx, pkID)
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
		Data:    planRecord,
	}
	utils.Render(w, response)
}
//SavePlan for inserting priceplans into table
func SavePlan(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	var plan models.Plans
	err := json.NewDecoder(r.Body).Decode(&plan)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	log.Infoft(ctx, "plan : %v ", plan)
	planData := plan
	planRecord, err := plans.SavePlan(ctx, planData)
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
		Description: "plan record is inserted successfully",
		Data:        planRecord,
	}
	utils.Render(w, response)
}
// UpdatePlan function is used to update the data in the plan table
func UpdatePlan(w http.ResponseWriter, r *http.Request) {
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
	var plan models.Plans
	if err := json.NewDecoder(r.Body).Decode(&plan); err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	rowsEffected, err := plans.UpdatePlan(ctx, pkID, &plan)
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
		Description: "plan record is Updated successfully",
		Data:        fmt.Sprintf("%d rows Effected ", rowsEffected),
	}
	utils.Render(w, response)
}
// DeletePlan for deleting plan record based on planId
func DeletePlan(w http.ResponseWriter, r *http.Request) {
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
		log.Infoft(ctx, "GetPlans : error %v ", err)
		return
	}
	rowsEffected, err := plans.DeletePlan(ctx, pkID)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	if rowsEffected != 0 {
		response := models.Response{
			Message:     utils.Success,
			Status:      http.StatusOK,
			Description: fmt.Sprintf(" %d rows deleted ", rowsEffected),
		}
		utils.Render(w, response)
	} else {
		response := models.Response{
			Error:       fmt.Sprintf(" %d ID should be in database ", rowsEffected),
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf(" %d rows deleted ", rowsEffected),
		}
		utils.Render(w, response)
	}
}
