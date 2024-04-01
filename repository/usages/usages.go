package repository

import (
	"fmt"
	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"log"
)

func SaveUsage(usageDtata models.Usages) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err.Error())
		return
	}
	fmt.Println("dbConnection : **************")
	ormDB := dbConnection.Create(&usageDtata)
	if ormDB.Error != nil {
		log.Println("Error inserting consumption:", ormDB.Error)
		return
	}
	fmt.Println("Inserted usages")
}
