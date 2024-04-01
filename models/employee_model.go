package models

type Employee struct {
	Id     int     `json:"id"`
	EmpId  int     `json:"empId"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}
