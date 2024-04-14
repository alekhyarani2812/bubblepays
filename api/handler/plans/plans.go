package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/plans"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

// Get plans Handler
func GetPlans(w http.ResponseWriter, r *http.Request) {
	plansList, err := plans.GetPlans()
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		log.Printf("GetPlans : error %v ", err)
		return
	}
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    plansList,
	}
	utils.Render(w, response)
	log.Println("consumptionList : success data ", plansList)
}

// Get plan by id Handler
func GetPlanById(w http.ResponseWriter, r *http.Request) {
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
	planRecord, err := plans.GetPlanById(pkId)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		log.Printf("GetPlandata : error %v ", err)
		return
	}
	log.Println("planRecord : ", planRecord)
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    planRecord,
	}
	utils.Render(w, response)
	log.Println("consumptionList : success data ", planRecord)
}

// Insert plans Handler
func SavePlan(w http.ResponseWriter, r *http.Request) {
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
	log.Println("plan : ", plan)
	planData := plan
	planRecord, err := plans.SavePlan(planData)
	log.Println("you are in handler after plans.SavePlan()", planRecord)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	log.Println("you are in SavesPlan handler out of err block")
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: "plan record is inserted successfully",
		Data:        planRecord,
	}
	utils.Render(w, response)
}

// Update plans Handler
func UpdatePlan(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdatePlan **********: ")
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
		log.Printf("GetPlans : error %v ", err)
		return
	}
	log.Println("pkId : ", pkId)
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
	log.Println("plan : ", plan)
	rowsEffected, err := plans.UpdatePlan(pkId, &plan)
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

// Delete plans Handler
func DeletePlan(w http.ResponseWriter, r *http.Request) {
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
		log.Printf("GetPlans : error %v ", err)
		return
	}
	rowsEffected, err := plans.DeletePlan(pkId)
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
		return
	} else {
		response := models.Response{
			Error:       fmt.Sprintf(" %d Id should be in database ", rowsEffected),
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf(" %d rows deleted ", rowsEffected),
		}
		log.Printf("***********failed to delete plan")
		utils.Render(w, response)
	}
}
