package api

import (
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	monthlyinvoice "github.com/bubblepays/repository/monthly_invoice"
	"github.com/bubblepays/utils"
)

// GetMonthlyInvoice godoc
// @Summary Get a list of MonthlyInvoices
// @Description Get a list of all MonthlyInvoice
// @Accept  json
// @Produce json
// @Param monthnum query string true "monthnum"
// @Success 200 {object} []models.MonthlyInvoice
// @Failure 400 {object} models.Response "Bad Request!"
// @Failure 404 {object} models.Response "Can not found monthnum"
// @Failure 401 {object} models.Response "Unauthorized Authentication is required"
// @Router /api/getmonthlyinvoice [get]
func GetMonthlyInvoice(w http.ResponseWriter, r *http.Request) {

	ctx := bcontext.WithTransactionID(r.Context())
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
		Description: "monthlyinvoice records are fetched successfully ",
		Data:        monthlyInvoiceData,
	}
	log.Info("from GetMonthlyInvoice handler")
	utils.Render(w, response)
}
