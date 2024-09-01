package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	planprice "github.com/bubblepays/repository/plan_price"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)
//LoadRoutes for routing all price api's
func LoadRoutes(router *mux.Router){
	router.HandleFunc("/api/getplanprice", GetPlanPrice).Methods("GET")
	router.HandleFunc("/api/getplanpricebyid/{id}", GetPlanPriceByID).Methods("GET")
	router.HandleFunc("/api/saveplanprice", SavePlanPrice).Methods("POST")
	router.HandleFunc("/api/updateplanpricebyid/{id}", UpdatePlanPrice).Methods("PUT")
	router.HandleFunc("/api/deleteplanpricebyid/{id}", DeletePlanPrice).Methods("DELETE")
}
// GetPlanPrice retrieves all planPrice records from the table
func GetPlanPrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	planpriceRecords, err := planprice.GetPlanPrice(ctx)
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
		Description: fmt.Sprintln("PlanPrice records we fetched successfully "),
		Data:        planpriceRecords,
	}
	utils.Render(w, response)
}
// GetPlanPriceByID Retrieve the planprice record based on the planpriceId
func GetPlanPriceByID(w http.ResponseWriter, r *http.Request) {
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
	planPriceRecord, err := planprice.GetPlanPriceByID(ctx, pkID)
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
		Description: fmt.Sprintf("PlanPrice record with id %d fetched successfully ", planPriceRecord.ID),
		Data:        planPriceRecord,
	}
	utils.Render(w, response)
}
// SavePlanPrice for inserting priceplans into table
func SavePlanPrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	var planPrice models.PlanPrice
	err := json.NewDecoder(r.Body).Decode(&planPrice)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	planPriceRecord, err := planprice.SavePlanPrice(ctx, &planPrice)
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
		Description: fmt.Sprintf("PlanPrice record with id %d inserted successfully ", planPriceRecord.ID),
		Data:        planPriceRecord,
	}
	utils.Render(w, response)
}
// UpdatePlanPrice function is used to update the data in the planprice table
func UpdatePlanPrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	strID := vars["id"]
	pkID, err := utils.StringToInt(strID)
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
	var planPrice models.PlanPrice
	erro := json.NewDecoder(r.Body).Decode(&planPrice)
	if erro != nil {
		response := models.Response{
			Error:   erro.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	planPriceRecord, err := planprice.UpdatePlanPrice(ctx, pkID, &planPrice)
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
		Data:    planPriceRecord,
	}
	utils.Render(w, response)
}
// DeletePlanPrice for deleting planprice record based on planpriceId
func DeletePlanPrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	varStr := mux.Vars(r)
	idStr := varStr["id"]
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
	rowsEffected, err := planprice.DeletePlanPrice(ctx, pkID)
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
		Description: fmt.Sprintf("PlanPrice record with id %d has been deleted", pkID),
	}
	utils.Render(w, response)
}
