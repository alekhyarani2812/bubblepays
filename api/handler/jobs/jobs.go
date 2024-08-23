package api

import (
	"net/http"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/repository/jobs"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

func LoadRoutes(router *mux.Router) {
	router.HandleFunc("/api/getjobs", GetJobs).Methods("GET")
}
func GetJobs(w http.ResponseWriter, r *http.Request) {
	ctx := bcontext.WithTransactionID(r.Context())
	jobList, err := jobs.GetJobs(ctx)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		log.Infoft(ctx, "GetJobs : error %v ", err)
		return
	}
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    jobList,
	}
	utils.Render(w, response)
}
