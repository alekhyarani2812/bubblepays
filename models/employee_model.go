package models

// Employee model
type Employee struct {
	ID     int     `json:"id"`
	EmpID  int     `json:"empId"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}
