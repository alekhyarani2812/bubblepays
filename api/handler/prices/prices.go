package api

import (
	"encoding/json"
	"fmt"
	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/prices"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetPrices(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello GetPrices handler:*********** ")
	priceReords, err := prices.GetPrices()
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
		//Description: fmt.Print("Succesfully you got prices data"),
		Data: priceReords,
	}
	utils.Render(w, response)
}

func GetPricesById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello GetPricesById handler:*********** ")
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
	priceReord, err := prices.GetPricesById(pkId)
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
		Status:  http.StatusFound,
		Data:    priceReord,
	}
	utils.Render(w, response)
}

func SavePrice(w http.ResponseWriter, r *http.Request) {
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
	priceRecord, err := prices.SavePrice(price)
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
	fmt.Println("Hello UpdatePrice handler:***********")
	vars := mux.Vars(r)
	idStr := vars["id"]
	pkId, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		log.Printf("GetPlans : error %v ", err)
		utils.Render(w, response)
		return
	}
	log.Println("pkId : ", pkId)
	var price models.Price
	err1 := json.NewDecoder(r.Body).Decode(&price)
	if err1 != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err1.Error(),
		}
		utils.Render(w, response)
		log.Printf("GetPlans : error %v ", err1)
		return
	}
	updatedRecord, err := prices.UpdatePrice(pkId, &price)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		log.Printf("GetPlans : error %v ", err)
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
	fmt.Println("Hello DeletePrice handler:***********")
	vars := mux.Vars(r)
	idStr := vars["id"]
	pkId, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		log.Printf("GetPlans : error %v ", err)
		utils.Render(w, response)
		return
	}
	log.Println("pkId : ", pkId)
	rowsEffected, err := prices.DeletePrice(pkId)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		log.Printf("GetPlans : error %v ", err)
		utils.Render(w, response)
		return
	}
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: fmt.Sprintf(" %d rows deleted from price table", rowsEffected),
	}
	log.Printf("GetPlans : error %v ", err)
	utils.Render(w, response)
}
