package plan_price

import (
	"context"
	"fmt"
	"log"

	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)

func GetPlanPrice(ctx context.Context) ([]models.PlanPrice, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	var planPrice []models.PlanPrice
	gorm := dbConnection.Find(&planPrice)
	if gorm.Error != nil {
		fmt.Println("error in getting records from planPrice ", err.Error())
		return nil, err
	}
	log.Println("planPrice :", planPrice)
	return planPrice, nil
}

func GetPlanPriceById(ctx context.Context, pkId int) (*models.PlanPrice, error) {
	fmt.Println("this is GetPlanPriceById repo")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	var planPrice models.PlanPrice
	if err := dbConnection.Where("id = ?", pkId).First(&planPrice).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("plan not found")
		} else {
			log.Printf("Error: %v\n", err)
		}
		return nil, err
	}
	log.Println("planPrice :", planPrice)
	return &planPrice, nil
}

func SavePlanPrice(ctx context.Context, planPriceData *models.PlanPrice) (*models.PlanPrice, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	gorm := dbConnection.Create(&planPriceData)
	if gorm.Error != nil && gorm != nil {
		fmt.Println("error in insertion of records", gorm.Error)
		return nil, err
	}
	fmt.Println("planPriceData :", planPriceData)
	return planPriceData, nil
}

func UpdatePlanPrice(ctx context.Context, pkId int, planPriceRecord *models.PlanPrice) (*models.PlanPrice, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	gorm := dbConnection.Model(&models.PlanPrice{}).Where("id = ?", pkId).Updates(&planPriceRecord)
	if gorm != nil && gorm.Error != nil {
		fmt.Println("error in updating data", err.Error())
		return nil, err
	}
	gorm.Find(&planPriceRecord)
	log.Println("planPriceRecord : ", planPriceRecord)
	return planPriceRecord, nil
}

func DeletePlanPrice(ctx context.Context, pkId int) (int64, error) {
	fmt.Println("this is DeletePlanPrice repo")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return 0, err
	}
	var planPrice models.PlanPrice
	gorm := dbConnection.Where("id = ?", pkId).Delete(&planPrice)
	if gorm.Error != nil && gorm != nil {
		fmt.Println("error in deleting PlanPrice data")
		return 0, err
	}
	rows := gorm.RowsAffected
	if rows == 0 {
		fmt.Println("error in deleting PlanPrice data ")
		return 0, err
	}
	fmt.Println("rows :", rows)
	return rows, nil
}
