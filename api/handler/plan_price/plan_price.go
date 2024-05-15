package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/models"
	planprice "github.com/bubblepays/repository/plan_price"
	"github.com/bubblepays/utils"
	"github.com/bubblepays/log"
	"github.com/gorilla/mux"
)

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

func GetPlanPriceById(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
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
	planPriceRecord, err := planprice.GetPlanPriceById(ctx, pkId)
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
		Description: fmt.Sprintf("PlanPrice record with id %d fetched successfully ", planPriceRecord.Id),
		Data:        planPriceRecord,
	}
	utils.Render(w, response)
}

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
		Description: fmt.Sprintf("PlanPrice record with id %d inserted successfully ", planPriceRecord.Id),
		Data:        planPriceRecord,
	}
	utils.Render(w, response)
}

func UpdatePlanPrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	vars := mux.Vars(r)
	strId := vars["id"]
	pkId, err := utils.StringToInt(strId)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	log.Infof("pkId : %v ", pkId)
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
	planPriceRecord, err := planprice.UpdatePlanPrice(ctx, pkId, &planPrice)
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

func DeletePlanPrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	varStr := mux.Vars(r)
	idStr := varStr["id"]
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
	log.Infof("pkId : %v ", pkId)
	rowsEffected, err := planprice.DeletePlanPrice(ctx, pkId)
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
		Description: fmt.Sprintf("PlanPrice record with id %d has been deleted", pkId),
	}
	utils.Render(w, response)
}
