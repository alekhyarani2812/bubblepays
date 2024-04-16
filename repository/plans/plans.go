package plans

import (
	"log"
	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)

func GetPlans() ([]models.Plans, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err)
		return nil, err
	}
	var plans []models.Plans
	if err := dbConnection.Find(&plans).Error; err != nil {
		log.Printf("failed to retrieve plans data: %v", err)
		return nil, err
	}
	return plans, nil
}

func GetPlanById(id int) (*models.Plans, error) {
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
	return &plan, nil
}

func SavePlan(planData models.Plans) (*models.Plans, error) {
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
	return &planData, nil
}

// Update the plan
func UpdatePlan(pkId int, plan *models.Plans) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return 0, err
	}
	gorm := dbConnection.Model(&models.Plans{}).Where("id = ?", pkId).Updates(&plan)
	rowsEffected := gorm.RowsAffected
	return rowsEffected, nil
}

func DeletePlan(pkId int) (int64, error) {
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
	return rowsEff, nil
}
