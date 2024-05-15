package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/prices"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

func GetPrices(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	priceReords, err := prices.GetPrices(ctx)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Status:  http.StatusBadRequest,
			Message: utils.Fail,
		}
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Status:  http.StatusFound,
		Message: utils.Success,
		Data:    priceReords,
	}
	utils.Render(w, response)
}

func GetPricesById(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
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
	priceReord, err := prices.GetPricesById(ctx, pkId)
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
		Status:  http.StatusFound,
		Data:    priceReord,
	}
	utils.Render(w, response)
}

func SavePrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	var price models.Price
	err := json.NewDecoder(r.Body).Decode(&price)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
		}
		utils.Render(w, response)
		return
	}
	priceRecord, err := prices.SavePrice(ctx, price)
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
		Status:      http.StatusCreated,
		Description: fmt.Sprintf("%d record inserted successfully", priceRecord.Id),
		Data:        priceRecord,
	}
	utils.Render(w, response)
}

func UpdatePrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
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
	log.Infoft(ctx, "pkId : %v ", pkId)
	var price models.Price
	err1 := json.NewDecoder(r.Body).Decode(&price)
	if err1 != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err1.Error(),
		}
		utils.Render(w, response)
		return
	}
	updatedRecord, err := prices.UpdatePrice(ctx, pkId, &price)
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
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: fmt.Sprintf("%d record updated successfully", pkId),
		Data:        updatedRecord,
	}
	utils.Render(w, response)
}

func DeletePrice(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
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
	log.Infoft(ctx, "pkId : %v ", pkId)
	rowsEffected, err := prices.DeletePrice(ctx, pkId)
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
		Description: fmt.Sprintf(" %d rows deleted from price table", rowsEffected),
	}
	utils.Render(w, response)
}
