package plans

import (
	"context"
	"fmt"
	"log"

	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"github.com/bubblepays/utils"
	"gorm.io/gorm"
)

func GetPlans(ctx context.Context) ([]models.Plans, error) {
	//get dbconnection from contex 
	dbConnection, err :=  utils.GetDBConnection(ctx)
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	var plans []models.Plans
	if err := dbConnection.Find(&plans).Error; err != nil {
		log.Printf("failed to retrieve plans data: %v", err)
		return nil, err
	}
	log.Println("plans : ", plans)
	return plans, nil
}

func GetPlanById(ctx context.Context, id int) (*models.Plans, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err)
		return nil, err
	}
	var plan models.Plans
	if err := dbConnection.Where("id = ?", id).First(&plan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("plan not found")
		} else {
			log.Printf("Error: %v\n", err)
		}
		return nil, err
	}
	log.Println("plan : ", plan)
	return &plan, nil
}

func SavePlan(ctx context.Context, planData models.Plans) (*models.Plans, error) {
	log.Println("planRecord", planData)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return nil, err
	}
	ormDB := dbConnection.Create(&planData)
	if ormDB.Error != nil {
		log.Println("Error inserting plan:", ormDB.Error)
		return nil, ormDB.Error
	}
	log.Println("planData : ", planData)
	return &planData, nil
}

// Update the plan
func UpdatePlan(ctx context.Context, pkId int, plan *models.Plans) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return 0, err
	}
	gorm := dbConnection.Model(&models.Plans{}).Where("id = ?", pkId).Updates(&plan)
	rowsEffected := gorm.RowsAffected
	log.Println("rowsEffected : ", rowsEffected)
	return rowsEffected, nil
}

func DeletePlan(ctx context.Context, pkId int) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return 0, err
	}
	var plan models.Plans
	result := dbConnection.Where("id = ?", pkId).Delete(&plan)
	if result.Error != nil {
		log.Println("***********failed to delete plan")
		return 0, result.Error
	}
	rowsEff := result.RowsAffected
	log.Println("rowsEff : ", rowsEff)
	return rowsEff, nil
}
