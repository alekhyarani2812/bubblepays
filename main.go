package main

import (
	"context"
	"net/http"

	authzApi "github.com/bubblepays/api/handler/authz"
	cardplanApi "github.com/bubblepays/api/handler/card_plan"
	consumptionApi "github.com/bubblepays/api/handler/consumption"
	employeeApi "github.com/bubblepays/api/handler/employee"
	monthlyinvoiceApi "github.com/bubblepays/api/handler/monthly_invoice"
	planpriceApi "github.com/bubblepays/api/handler/plan_price"
	planApi "github.com/bubblepays/api/handler/plans"
	priceApi "github.com/bubblepays/api/handler/prices"
	serviceApi "github.com/bubblepays/api/handler/service"
	serviceenrollmentApi "github.com/bubblepays/api/handler/service_enrollment"
	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/database"
	_ "github.com/bubblepays/docs" // This is where Swagger comments are located
	"github.com/bubblepays/jobs"
	"github.com/bubblepays/log"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dbConnection, dbErr := database.GetORMDbConnection()
		if dbErr != nil {
			log.Infof("err %v ", dbErr.Error())
			return
		}
		log.Info("Middleware executing")
		transactionID := bcontext.GenerateTransactionID()
		TransactionsIDKey := bcontext.TransactionIDKey
		// Create a new context with the transaction ID
		ctx := context.WithValue(r.Context(), TransactionsIDKey, transactionID)
		//set dbconnection in context
		ctx = context.WithValue(ctx, bcontext.DBconnectionKey, dbConnection)
		// Serve the request with the new context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
type Emp struct{
	Name string
}
func UpdateEmp(e Emp) Emp{
	e.Name = "Alekhya"
	return e
}
// @title Bubblepays API
// @version 1.0
// @description This is a sample API
// @host localhost:8080
// @BasePath /
func main() {
	e := Emp{Name : "Satya"}
	e = UpdateEmp(e)
	log.Info(e)
	dbConnection, dbErr := database.GetORMDbConnection()
	if dbErr != nil {
		log.Errorf("err %v ", dbErr.Error())
		return
	}
	go jobs.RunMachineJob(dbConnection)
	router := mux.NewRouter()
	//add middleware
	router.Use(middleware)
	loadAllRoutes(router)
	//jobs.RunMachineJob()
	// Swagger UI route
	router.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The Swagger JSON endpoint
	))
	//To start the port with 8080
	http.Handle("/", router)
	log.Info("Started server on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Errorf(" ERROR ListenAndServe err: %v", err)
	}
}
func loadAllRoutes(router *mux.Router){
	authzApi.LoadRoutes(router)
	consumptionApi.LoadRoutes(router)
	monthlyinvoiceApi.LoadRoutes(router)
	cardplanApi.LoadRoutes(router)
	planpriceApi.LoadRoutes(router)
	planApi.LoadRoutes(router)
	priceApi.LoadRoutes(router)
	serviceApi.LoadRoutes(router)
	serviceenrollmentApi.LoadRoutes(router)
	employeeApi.LoadRoutes(router)
}
