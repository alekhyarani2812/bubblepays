package api

import (
	"log"
	"net/http"
	"todoproject/database"
	"todoproject/models"
	"todoproject/utils"
	_ "github.com/go-sql-driver/mysql"
)

func GetEmployees(w http.ResponseWriter) []models.Employee {
	var empArray []models.Employee
	db := database.GetDbConnection()
	defer db.Close()
	//Perform a SELECT query
	rows, err := db.Query("SELECT * FROM employee")
	if err != nil {
		response := models.Response{
			Message: utils.Fail,
			Status:  http.StatusInternalServerError,
			Error:   err.Error(),
		}
		utils.Render(w, response)
	}
	defer rows.Close()
	log.Println("select query compiled")
	// Iterate over the result set
	for rows.Next() {
		var employee models.Employee
		if err := rows.Scan(&employee.Id, &employee.EmpId, &employee.Name, &employee.Age, &employee.Salary); err != nil {
			log.Fatal(err)
		}
		empArray = append(empArray, employee)
	}
	log.Println("***********", empArray)
	return empArray
}
// Get employees request
func GetEmployeesJsonData(w http.ResponseWriter, r *http.Request) {
	empList := GetEmployees(w)
	response := models.Response{
		Message: utils.Success,
		Status:  http.StatusOK,
		Data:    empList,
	}
	utils.Render(w, response)
}
