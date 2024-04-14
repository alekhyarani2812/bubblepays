package cardplan

import (
	"fmt"

	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
)

func GetCardPlan() ([]models.CardPlan, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	var cardplan []models.CardPlan
	gorm := dbConnection.Find(&cardplan)
	if gorm.Error != nil && gorm != nil {
		errors := gorm.Error
		fmt.Println("error in getting records from cardplan ", errors)
		return nil, errors
	}
	fmt.Println("cardplan : ", cardplan)
	return cardplan, nil
}

func GetCardPlanById(pkId int) (*models.CardPlan, error){
	fmt.Println("GetCardPlanById repo ************")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	var cardPlanRecord models.CardPlan
	gorm := dbConnection.Where("id = ?",pkId).Find(&cardPlanRecord)
	if gorm != nil && gorm.Error != nil {
		fmt.Println("error in getting CardPlan data")
		return nil, err
	}
	fmt.Println(cardPlanRecord)
	return &cardPlanRecord, nil
}

func SaveCardPlan(cardPlanData *models.CardPlan) (*models.CardPlan, error) {
	fmt.Println("this is CardPlan repo  CardPlan : ", cardPlanData)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ", err.Error())
		return nil, err
	}
	gorm := dbConnection.Create(&cardPlanData)
	if gorm.Error != nil && gorm != nil {
		fmt.Println("error in insertion of records", gorm.Error)
		return nil, err
	}
	fmt.Println("cardPlanData : ", cardPlanData)
	return cardPlanData, nil
}

func UpdateCardPlan(pkId int, cardPlanRecord *models.CardPlan) (*models.CardPlan, error){
	fmt.Println("this is cardplan repo cardPlanRecord :  ,pkId", cardPlanRecord, pkId)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ")
		return nil, err
	}
	//cardPlanRecord
	gorm := dbConnection.Model(&models.CardPlan{}).Where("id = ?", pkId).Updates(&cardPlanRecord)
	if gorm != nil && gorm.Error != nil {
		fmt.Println("error in updating cardPlan Record ")
		return nil, err
	}
	gorm.Find(&cardPlanRecord)
	fmt.Println("cardPlanRecord : ",cardPlanRecord)
	return cardPlanRecord, nil
}

func DeleteCardPlan(pkId int) (int64, error){
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connection ")
		return 0, err
	}
	var cardPlanRecord models.CardPlan
	gorm := dbConnection.Where("id = ?", pkId).Delete(&cardPlanRecord)
	if gorm.Error != nil {
		fmt.Println("error in deleting card PlanRecord *****gorm.Error != nil*****")
		return 0, err
	}
	rows := gorm.RowsAffected
	if rows == 0 {
		fmt.Println("error in deleting cardPlan Record")
		return 0, err
	}
	fmt.Println(" rows effected : ", rows)
	return rows, nil
}
