package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bubblepays/database"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

func GetEmployees(w http.ResponseWriter) ([]models.Employee, error) {
	var empArray []models.Employee
	db := database.GetDbConnection()
	defer db.Close()
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return nil, err
	}
	defer rows.Close()
	log.Info("select query compiled")
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.Id, &employee.EmpId, &employee.Name, &employee.Age,
			&employee.Salary); err != nil {
			log.Fatal(err.Error())
		}
		empArray = append(empArray, employee)
	}
	log.Infof("*****empArray****** %v ", empArray)
	return empArray, nil
}

func GetEmployeesJsonData(w http.ResponseWriter, r *http.Request) {
	empList, err := GetEmployees(w)
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
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    empList,
	}
	utils.Render(w, response)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	db := database.GetDbConnection()
	defer db.Close()
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
	log.Infof(" Employee iD %v ", pkId)
	var employee models.Employee
	rows, err := db.Query("select * from employee where id = ?", pkId)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	for rows.Next() {
		if err := rows.Scan(&employee.Id, &employee.EmpId, &employee.Name, &employee.Age,
			&employee.Salary); err != nil {
			response := models.Response{
				Message: utils.Fail,
				Status:  http.StatusBadRequest,
				Error:   err.Error(),
			}
			utils.Render(w, response)
			return
		}
	}
	successResponse := models.Response{
		Message: utils.Success,
		Data:    employee,
		Status:  http.StatusCreated,
	}
	utils.Render(w, successResponse)
	log.Infof("employee : %v", employee)
}

func SaveEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee
	w.Header().Set("Content-Type", "application/json")
	db := database.GetDbConnection()
	defer db.Close()
	// Decode request body into Employee struct
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errorMessage := err.Error()
		errorResponse := models.Response{
			Error: errorMessage,
		}
		json.NewEncoder(w).Encode(errorResponse)
		return
	}
	_, err := db.Exec("INSERT INTO employee (empId, name, age, salary) VALUES (?,?,?,?)",
		employee.EmpId, employee.Name, employee.Age, employee.Salary)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	description := fmt.Sprintf("Employee '%s' record is inserted successfully", employee.Name)
	successResponse := models.Response{
		Message:     utils.Success,
		Description: description,
		Data:        employee,
		Status:      http.StatusCreated,
	}
	utils.Render(w, successResponse)
	log.Infof(" Employee %v", employee.Name, " record is inserted")
}

func UpdateEmpHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// idStr := vars["id"]
	idStr := r.PathValue("id")
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
	log.Infof("pkId is %v ", pkId)

	var employee models.Employee
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
	result, err := db.Exec("UPDATE employee SET empId = ?, name=?, age = ?, salary=? WHERE id = ?",
		employee.EmpId, employee.Name, employee.Age, employee.Salary, pkId)
	if err != nil {
		response := models.Response{
			Error:  err.Error(),
			Status: http.StatusInternalServerError,
		}
		utils.Render(w, response)
		log.Info("update Handler  result, err")
		return
	}
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
	if rowsAffected != 0 {
		response := models.Response{
			Message: "employee record is Updated in successfully",
			Status:  http.StatusOK,
			Data:    employee,
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
	log.Infof("update Handler %v ", rowsAffected)
}

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
			Message: utils.Success,
			Status:  http.StatusOK,
			Description: fmt.Sprintf("Id = '%v' record is deleted successfully %v row is effected",
				id, rowsAffected),
		}
		utils.Render(w, response)
	}
}
