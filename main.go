package main

import (
	"fmt"
	"log"
	"net/http"

	consumptionApi "github.com/bubblepays/api/handler/consumption"
	employeeApi "github.com/bubblepays/api/handler/employee"
	planApi "github.com/bubblepays/api/handler/plans"
	serviceApi "github.com/bubblepays/api/handler/service"
	"github.com/bubblepays/database"
	"github.com/bubblepays/jobs"

	"github.com/gorilla/mux"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform middleware tasks here, e.g., logging
		fmt.Println("Middleware executing")
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("This is main package")
	//learning.Implimentation()
	//create a job
	dbConnection, dbErr := database.GetORMDbConnection()
	if dbErr != nil {
		log.Printf("err %v ", dbErr.Error())
		return
	}
	go jobs.RunMachineJob(dbConnection)

	router := mux.NewRouter()

	//add middleware
	router.Use(middleware)

	router.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}).Methods("GET")

	// attach helloworld handler
	// getMessageHandler := middleware(http.HandlerFunc(api.GetMessageHandler))
	// mux.Handle("/api/message", getMessageHandler)

	//Plans
	router.HandleFunc("/api/getplans", planApi.GetPlans).Methods("GET")
	router.HandleFunc("/api/getplanbyid/{id}", planApi.GetPlanById).Methods("GET")
	router.HandleFunc("/api/saveplan", planApi.SavePlan).Methods("POST")
	router.HandleFunc("/api/updateplan/{id}", planApi.UpdatePlan).Methods("PUT")
	router.HandleFunc("/api/deleteplan/{id}", planApi.DeletePlan).Methods("DELETE")

	//Service
	router.HandleFunc("/api/getservices", serviceApi.GetServices).Methods("GET")
	router.HandleFunc("/api/saveservice", serviceApi.SaveService).Methods("POST")
	router.HandleFunc("/api/deleteservice/{id}", serviceApi.DeleteServiceById).Methods("DELETE")

	//consumption
	router.HandleFunc("/api/consumptions", consumptionApi.GetConsumptions).Methods("GET")
	router.HandleFunc("/api/consumptions/{id}", consumptionApi.GetConsumptionById).Methods("GET")
	router.HandleFunc("/api/saveconsumption", consumptionApi.SavesConsumption).Methods("POST")
	router.HandleFunc("/api/deleteconsumption/{id}", consumptionApi.DeleteConsumption).Methods("DELETE")

	//Employee api's
	router.HandleFunc("/api/employees", employeeApi.GetEmployeesJsonData).Methods("GET")
	router.HandleFunc("/api/employee", employeeApi.SaveEmployee).Methods("POST")
	router.HandleFunc("/api/deleteEmployee/{id}", employeeApi.DeleteEmpHandler).Methods("DELETE")
	router.HandleFunc("/api/updateEmployee/{id}", employeeApi.UpdateEmpHandler).Methods("PUT")

	//jobs.RunMachineJob()

	http.Handle("/", router)
	fmt.Println("Started server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf(" ERROR ListenAndServe err: %v", err)
	}
	//mypkg.TestFunction()
	//employee.Printempdata()
	//api.ApiCalling()
	//api.EmployeeApiCalling()
	//context.
	//Start()
	//getMessage()
	//GoRoutine()
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler executing")
	w.Write([]byte("Hello, World!"))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler executing2")
	w.Write([]byte("Hello, World2!"))
}
