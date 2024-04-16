package main

import (
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
	"github.com/gorilla/mux"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Middleware executing")
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	fmt.Println("This is main package")
	dbConnection, dbErr := database.GetORMDbConnection()
	if dbErr != nil {
		log.Printf("err %v ", dbErr.Error())
		return
	}
	go jobs.RunMachineJob(dbConnection)

	router := mux.NewRouter()

	//add middleware
	router.Use(middleware)
	// attach helloworld handler
	// getMessageHandler := middleware(http.HandlerFunc(api.GetMessageHandler))
	// mux.Handle("/api/message", getMessageHandler)

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

	http.Handle("/", router)
	fmt.Println("Started server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf(" ERROR ListenAndServe err: %v", err)
	}
	router.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	}).Methods("GET")
}
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler executing")
	w.Write([]byte("Hello, World!"))
}

func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler executing2")
	w.Write([]byte("Hello, World2!"))
}
