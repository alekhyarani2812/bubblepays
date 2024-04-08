package plans

import (
	"log"

	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)

// Get all the Plans
func GetPlans() ([]models.Plans, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err)
		return nil, err
	}
	var plans []models.Plans
	if err := dbConnection.Find(&plans).Error; err != nil {
		log.Printf("failed to retrieve consumption data: %v", err)
		return nil, err
	}
	log.Printf("SIZE of consumption: size:: %v ", len(plans))
	return plans, nil
}

// Get plans by id
func GetPlanById(id int) ([]models.Plans, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err)
		return nil, err
	}
	var plan []models.Plans
	if err := dbConnection.Where("id = ?", id).First(&plan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("plan not found")
		} else {
			log.Printf("Error: %v\n", err)
		}
		return nil, err
	}
	log.Printf("plan found: %+v\n", plan)
	log.Printf("SIZE of plan: size:: %v ", len(plan))
	return plan, nil
}

// Insert plan into db
func SavePlan(planData models.Plans) (*models.Plans, error) {
	log.Println("planRecord", planData)
	// get orm db connection
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
	log.Println("plan inserted successfully")
	log.Println("ormDB  :  ", planData)
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
	log.Printf("rowsEffected : %+v\n", rowsEffected)
	return rowsEffected, nil
}

// Delete the plan
func DeletePlan(pkId int) (int64, error){
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Println("err in dbConnection :****")
		return 0,err
	}
	var plan models.Plans
	result := dbConnection.Where("id = ?", pkId).Delete(&plan)
	if result.Error != nil {
		log.Println("failed to delete consumption")
	}
	rowsEff := result.RowsAffected
	return rowsEff, nil
}
