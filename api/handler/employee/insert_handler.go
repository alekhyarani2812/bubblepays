package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
)

// Post request employee insert record
func SaveEmployee(w http.ResponseWriter, r *http.Request) {
	var employee models.Employee

	//for json response
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
	_, err := db.Exec("INSERT INTO employee (empId, name, age, salary) VALUES (?,?,?,?)", employee.EmpId, employee.Name, employee.Age, employee.Salary)
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

	empRecord, err := getEmployeeByEmployeeID(db, employee.EmpId)
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
		return
	}
	fmt.Println(empRecord)
	successResponse := models.Response{
		Message:     utils.Success,
		Description: description,
		Data:        empRecord,
		Status:      http.StatusCreated,
	}
	utils.Render(w, successResponse)
	log.Println(" Employee ", employee.Name, " record is inserted")
}

func getEmployeeByEmployeeID(db *sql.DB, empId int) (*models.Employee, error) {

	var employee models.Employee
	//var empArray []models.Employee
	rows, err := db.Query("select * from employee where empId =?", empId)
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
	return &employee, nil
}
