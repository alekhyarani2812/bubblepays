package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"todoproject/models"
	"todoproject/repository/consumption"
	"todoproject/utils"
)

// Get all consumptions
func GetConsumptions(w http.ResponseWriter, r *http.Request) {
	consumptionList, err := consumption.GetConsumptions()
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		log.Printf("GetConsumptions : error %v ", err)
		return
	}
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    consumptionList,
	}
	utils.Render(w, response)
	log.Println("consumptionList : success data ", consumptionList)
}

// Get 1 employee data by Id primarykey
func GetConsumptionById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	//idStr := r.PathValue("id")
	pkId, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return

	}
	log.Println("pkId is", pkId)
	consumptionList, err := consumption.GetConsumptionById(pkId)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		log.Printf("GetConsumptions : error %v ", err)
		return
	}
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    consumptionList,
	}
	utils.Render(w, response)
	log.Println("consumptionList : success data ", consumptionList)
}

// Insert consumption data
func SavesConsumption(w http.ResponseWriter, r *http.Request) {
	var consum models.Consumption
	err := json.NewDecoder(r.Body).Decode(&consum)
	if err != nil {
		response := models.Response{
			Error:   err.Error(),
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	log.Println("consum : ", consum)
	consumData := consum
	consumptionRecord, err := consumption.SavesConsumption(consumData)
	log.Println("you are in handler after consumption.SavesConsumption()")
	if err != nil {
		log.Println("you are in SavesConsumption handler err block")
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	log.Println("you are in SavesConsumption handler out of err block")
	response := models.Response{
		Message:     utils.Success,
		Status:      http.StatusOK,
		Description: "consumption record is inserted successfully",
		Data:        consumptionRecord,
	}
	utils.Render(w, response)
	log.Println("consumptionList : success data ", consumptionRecord)
}

func DeleteConsumption(w http.ResponseWriter, r *http.Request) {
	log.Println("this is DeleteConsumption Handler")
	vars := mux.Vars(r)
	idStr := vars["id"]
	//idStr := r.PathValue("id")
	pkId, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	log.Println("pkId is", pkId)
	rowsEffected, err := consumption.DeleteConsumption(pkId)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	log.Println("rows effected : ", rowsEffected)
	if rowsEffected != 0 {
		response := models.Response{
			Message:     utils.Success,
			Status:      http.StatusOK,
			Description: fmt.Sprintf("consumption record %v is deleted", pkId),
		}
		utils.Render(w, response)
		fmt.Println("consumption with deleted successfully", pkId)
		return
	} else {
		response := models.Response{
			Message:     utils.Fail,
			Status:      http.StatusBadRequest,
			Description: fmt.Sprintf("consumption record %v does not exist in consumption table", pkId),
		}
		utils.Render(w, response)
	}
}

//service talble
//Id pk, serviceId alphanumeric, active boolean, location
