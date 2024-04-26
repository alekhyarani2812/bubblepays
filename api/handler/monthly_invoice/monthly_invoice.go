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
	monthlyInvoiceData, err := monthlyinvoice.GetMonthlyInvoice(ctx)
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
		Data:        monthlyInvoiceData ,
	}
	fmt.Println("from GetMonthlyInvoice handler")
	utils.Render(w, response)
}
