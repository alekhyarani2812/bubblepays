package main

// @title My Go API
// @version 1.0
// @description This is a sample Go API with Swagger documentation.
// @BasePath /api/v1

import (
	"context"
	"fmt"
	"log"
	"net/http"

	cardplanApi "github.com/bubblepays/api/handler/card_plan"
	consumptionApi "github.com/bubblepays/api/handler/consumption"
	employeeApi "github.com/bubblepays/api/handler/employee"
	planpriceApi "github.com/bubblepays/api/handler/plan_price"
	planApi "github.com/bubblepays/api/handler/plans"
	priceApi "github.com/bubblepays/api/handler/prices"
	serviceApi "github.com/bubblepays/api/handler/service"
	serviceenrollmentApi "github.com/bubblepays/api/handler/service_enrollment"
	"github.com/bubblepays/database"
	"github.com/bubblepays/jobs"
	"github.com/bubblepays/utils"
	"github.com/gorilla/mux"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		dbConnection, dbErr := database.GetORMDbConnection()
		if dbErr != nil {
			log.Printf("err %v ", dbErr.Error())
			return
		}
		fmt.Println("Middleware executing")
		transactionID := utils.GenerateTransactionID()
		TransactionsIDKey := utils.TransactionIDKey
		//var myConnection string
		// Create a new context with the transaction ID
		ctx := context.WithValue(r.Context(), TransactionsIDKey, transactionID)
		//ctx.Value(dbConnection)
		ctx = context.WithValue(r.Context(), utils.DBconnectionKey, dbConnection)
		
		// Serve the request with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func main() {
	fmt.Println("This is main package")
	//fmt.Println(os.Args)

	dbConnection, dbErr := database.GetORMDbConnection()
	if dbErr != nil {
		log.Printf("err %v ", dbErr.Error())
		return
	}
	go jobs.RunMachineJob(dbConnection)
	router := mux.NewRouter()
	//add middleware
	router.Use(middleware)

	//CardPlan
	router.HandleFunc("/api/getcardplan", cardplanApi.GetCardPlan).Methods("GET")
	router.HandleFunc("/api/getcardplanbyid/{id}", cardplanApi.GetCardPlanById).Methods("GET")
	router.HandleFunc("/api/savecardplan", cardplanApi.SaveCardPlan).Methods("POST")
	router.HandleFunc("/api/updatecardplan/{id}", cardplanApi.UpdateCardPlan).Methods("PUT")
	router.HandleFunc("/api/deletecardplan/{id}", cardplanApi.DeleteCardPlan).Methods("DELETE")

	//plan_price
	router.HandleFunc("/api/getplanprice", planpriceApi.GetPlanPrice).Methods("GET")
	router.HandleFunc("/api/getplanpricebyid/{id}", planpriceApi.GetPlanPriceById).Methods("GET")
	router.HandleFunc("/api/saveplanprice", planpriceApi.SavePlanPrice).Methods("POST")
	router.HandleFunc("/api/updateplanpricebyid/{id}", planpriceApi.UpdatePlanPrice).Methods("PUT")
	router.HandleFunc("/api/deleteplanpricebyid/{id}", planpriceApi.DeletePlanPrice).Methods("DELETE")

	//ServiceEnrollment
	router.HandleFunc("/api/getserviceenrollment", serviceenrollmentApi.GetServiceEnrollment).Methods("GET")
	router.HandleFunc("/api/getserviceenrollmentbyid/{id}", serviceenrollmentApi.GetServiceEnrollmentById).Methods("GET")
	router.HandleFunc("/api/saveserviceenrollment", serviceenrollmentApi.SaveServiceEnrollment).Methods("POST")
	router.HandleFunc("/api/updateserviceenrollment/{id}", serviceenrollmentApi.UpdateServiceEnrollment).Methods("PUT")
	router.HandleFunc("/api/deleteserviceenrollment/{id}", serviceenrollmentApi.DeleteServiceEnrollment).Methods("DELETE")

	//prices
	router.HandleFunc("/api/getprices", priceApi.GetPrices).Methods("GET")
	router.HandleFunc("/api/getpricesbyid/{id}", priceApi.GetPricesById).Methods("GET")
	router.HandleFunc("/api/saveprice", priceApi.SavePrice).Methods("POST")
	router.HandleFunc("/api/updateprice/{id}", priceApi.UpdatePrice).Methods("PUT")
	router.HandleFunc("/api/deleteprice/{id}", priceApi.DeletePrice).Methods("DELETE")

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
	router.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		// Get the context from the request
		ctx := r.Context()
		w.Write([]byte("Hello transactionId: " + utils.GetTransactionID(ctx)))
	}).Methods("GET")

	//To start the port with 8080
	http.Handle("/", router)
	fmt.Println("Started server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf(" ERROR ListenAndServe err: %v", err)
	}
}
func handler(w http.ResponseWriter, r *http.Request) {
	//jdbnjfblh vn
	fmt.Println("Handler executing")
	w.Write([]byte("Hello, World!"))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler executing2")
	w.Write([]byte("Hello, World2!"))
}
