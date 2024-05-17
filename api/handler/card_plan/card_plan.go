package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/models"
	cardplan "github.com/bubblepays/repository/card_plan"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router){
	router.HandleFunc("/api/getcardplan", GetCardPlan).Methods("GET")
	router.HandleFunc("/api/getcardplanbyid/{id}", GetCardPlanByID).Methods("GET")
	router.HandleFunc("/api/savecardplan", SaveCardPlan).Methods("POST")
	router.HandleFunc("/api/updatecardplan/{id}", UpdateCardPlan).Methods("PUT")
	router.HandleFunc("/api/deletecardplan/{id}", DeleteCardPlan).Methods("DELETE")
}

func GetCardPlan(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	cardPlanRecords, err := cardplan.GetCardPlan(ctx)
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

func GetCardPlanByID(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	varStr := mux.Vars(r)
	idStr := varStr["id"]
	pkID, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
		}
		utils.Render(w, response)
		return
	}
	cardPlanRecord, err := cardplan.GetCardPlanByID(ctx, pkID)
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
	ctx := bcontext.WithTransactionID(r.Context())
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
	cardPlanRecord, err := cardplan.SaveCardPlan(ctx, &cardPlan)
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
	cardPlanRecordData, err := cardplan.UpdateCardPlan(ctx, pkID, &cardPlanRecord)
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
		Description: fmt.Sprintf("record with id %d has been updated", pkID),
		Data:        cardPlanRecordData,
	}
	utils.Render(w, response)
}

func DeleteCardPlan(w http.ResponseWriter, r *http.Request) {
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
	}
	rowsEffected, err := cardplan.DeleteCardPlan(ctx, pkID)
	if err != nil {
		response := models.Response{
			Error:       err.Error(),
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf("you should err.Error()%d", pkID),
		}
		utils.Render(w, response)
		return
	}
	if pkID == 0 {
		response := models.Response{
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf("you should not enter id is : %d : ", pkID),
		}
		utils.Render(w, response)
		return
	}
	if rowsEffected != 0 {
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
		Description: fmt.Sprintf("cardplan record with id %d has been deleted", pkID),
	}
	utils.Render(w, response)
}
