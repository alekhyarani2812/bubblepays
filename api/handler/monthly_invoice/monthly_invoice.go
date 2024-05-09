package api

import (
	"fmt"
	"net/http"

	"github.com/bubblepays/models"
	monthlyinvoice "github.com/bubblepays/repository/monthly_invoice"
	"github.com/bubblepays/utils"
)

func GetMonthlyInvoice(w http.ResponseWriter, r *http.Request) {
	ctx := utils.WithTransactionID(r.Context())
	queryParams := r.URL.Query()
	monthNumQparam := queryParams.Get("monthnum")
	if monthNumQparam == "" {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return	
	}
	monthlyInvoiceData, err := monthlyinvoice.GetMonthlyInvoice(ctx, monthNumQparam)
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
		Description: fmt.Sprintln("monthlyinvoice records are fetched successfully "),
		Data:        monthlyInvoiceData,
	}
	fmt.Println("from GetMonthlyInvoice handler")
	utils.Render(w, response)
}

func GetMonthlyInvoiceById(w http.ResponseWriter, r *http.Request){
	ctx := utils.WithTransactionID(r.Context())
	fmt.Println("this is GetMonthlyInvoiceById handler")
	monthlyinvoice.GetMonthlyInvoiceById(ctx)
}