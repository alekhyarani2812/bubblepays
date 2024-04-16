package api

import (
	"fmt"
	"net/http"

	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"

	"github.com/gorilla/mux"
)

// delete employee rows
func DeleteEmpHandler(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
	idStr := vars["id"]
	db := database.GetDbConnection()
	defer db.Close()

	id, err := utils.StringToInt(idStr)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	result, err := db.Exec("DELETE FROM employee WHERE id=?", id)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	if rowsAffected == 0 {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   fmt.Sprintf("Id = %v does not exist in the Employee table", idStr),
		}
		utils.Render(w, response)
	} else {
		response := models.Response{
			Message:     utils.Success,
			Status:      http.StatusOK,
			Description: fmt.Sprintf("Id = '%v' record is deleted successfully %v row is effected ", id, rowsAffected),
		}
		utils.Render(w, response)
	}
}
