package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"todoproject/database"
	"todoproject/models"
	"todoproject/utils"
)

func UpdateEmpHandler(w http.ResponseWriter, r *http.Request) {
	//idStr := r.URL.Query().Get("id")
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
	fmt.Println("pkId is", pkId)

	var employee models.Employee
	//get db connection
	db := database.GetDbConnection()
	defer db.Close()

	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		response := models.Response{
			Error:   "Invalid request body",
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
		return
	}
	//update employee records
	result, err := db.Exec("UPDATE employee SET empId = ?, name = ?, age = ?, salary = ? WHERE id = ?", employee.EmpId, employee.Name, employee.Age, employee.Salary, pkId)
	if err != nil {
		response := models.Response{
			Error:  err.Error(),
			Status: http.StatusInternalServerError,
		}
		utils.Render(w, response)
		fmt.Println("update Handler  result, err")
		return
	}
	// Check if any rows were affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Error:   err.Error(),
			Status:  http.StatusInternalServerError,
		}
		utils.Render(w, response)
		return
	}
	empRecord, err := getEmployeeByID(db, pkId)
	fmt.Println("empRecord : ", empRecord)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	if rowsAffected != 0 {
		response := models.Response{
			Message: "employee record is Updated in successfully",
			Status:  http.StatusOK,
			Data:    empRecord,
		}
		utils.Render(w, response)
	} else {
		response := models.Response{
			Message: utils.Fail,
			Error:   "No rows effected",
			Status:  http.StatusBadRequest,
		}
		utils.Render(w, response)
	}
	fmt.Println("update Handler", rowsAffected)
}

func getEmployeeByID(db *sql.DB, id int) (*models.Employee, error) {

	var employee models.Employee
	//var empArray []models.Employee
	rows, err := db.Query("select * from employee where id = ?", id)
	if err != nil {
		fmt.Println("err : ", err)
		return nil, err
	}

	for rows.Next() {
		if err := rows.Scan(&employee.Id, &employee.EmpId, &employee.Name, &employee.Age, &employee.Salary); err != nil {
			return nil, err
		}
		fmt.Println("employee : ", employee)
	}
	fmt.Println("employee : ", employee)
	if employee.Id == 0 {
		return nil, nil
	}
	return &employee, nil
}
