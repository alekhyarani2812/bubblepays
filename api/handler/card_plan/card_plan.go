package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bubblepays/models"
	cardplan "github.com/bubblepays/repository/card_plan"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

func GetCardPlan(w http.ResponseWriter, r *http.Request) {
	cardPlanRecords, err := cardplan.GetCardPlan()
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
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    cardPlanRecords,
	}
	utils.Render(w, response)
}

func GetCardPlanById(w http.ResponseWriter, r *http.Request) {
	varStr := mux.Vars(r)
	idStr := varStr["id"]
	pkId, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
		}
		utils.Render(w, response)
		return
	}
	fmt.Println("GetCardPlanById Handler ******pkId: ******", pkId)
	cardPlanRecord, err := cardplan.GetCardPlanById(pkId)
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
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    cardPlanRecord,
	}
	utils.Render(w, response)
}

func SaveCardPlan(w http.ResponseWriter, r *http.Request) {
	var cardPlan models.CardPlan
	err := json.NewDecoder(r.Body).Decode(&cardPlan)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	cardPlanRecord, err := cardplan.SaveCardPlan(&cardPlan)
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
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    cardPlanRecord,
	}
	utils.Render(w, response)
}

func UpdateCardPlan(w http.ResponseWriter, r *http.Request) {
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
	}
	var cardPlanRecord models.CardPlan
	erro := json.NewDecoder(r.Body).Decode(&cardPlanRecord)
	if erro != nil {
		response := models.Response{
			Error:   erro.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
	}
	cardPlanRecordData, err := cardplan.UpdateCardPlan(pkId, &cardPlanRecord)
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
		Status:      http.StatusOK,
		Description: fmt.Sprintf("record with id %d has been updated", pkId),
		Data:        cardPlanRecordData,
	}
	utils.Render(w, response)
}

func DeleteCardPlan(w http.ResponseWriter, r *http.Request) {
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
	}
	rowsEffected, err := cardplan.DeleteCardPlan(pkId)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Description: fmt.Sprintf("you should err.Error()%d", pkId),
		}
		utils.Render(w, response)
		return
	}
	if pkId == 0 {
		response := models.Response{
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf("you should not enter id is : %d : ", pkId),
		}
		utils.Render(w, response)
		return
	}
	if rowsEffected != 0{
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
		Description: fmt.Sprintf("cardplan record with id %d has been deleted", pkId),
	}
	utils.Render(w, response)
}
