package plans

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/database"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)

func GetPlans(ctx context.Context) ([]models.Plans, error) {
	//get dbconnection from contex
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	var plans []models.Plans
	if err := dbConnection.Find(&plans).Error; err != nil {
		log.Infof("failed to retrieve plans data: %v", err)
		return nil, err
	}
	log.Infof("plans : %v ", plans)
	return plans, nil
}

func GetPlanByID(ctx context.Context, id int) (*models.Plans, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("err %v ", err)
		return nil, err
	}
	var plan models.Plans
	if err := dbConnection.Where("id = ?", id).First(&plan).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Info("plan not found")
		} else {
			log.Infof("Error: %v\n", err)
		}
		return nil, err
	}
	log.Infof("plan : %v ", plan)
	return &plan, nil
}

func SavePlan(ctx context.Context, planData models.Plans) (*models.Plans, error) {
	log.Infof("planRecord %v", planData)
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Info("err in dbConnection :****")
		return nil, err
	}
	ormDB := dbConnection.Create(&planData)
	if ormDB.Error != nil {
		log.Infof("Error inserting plan: %v ", ormDB.Error)
		return nil, ormDB.Error
	}
	log.Infof("planData : %v ", planData)
	return &planData, nil
}

//UpdatePlan Update the plan
func UpdatePlan(ctx context.Context, pkID int, plan *models.Plans) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("err in dbConnection :****")
		return 0, err
	}
	gorm := dbConnection.Model(&models.Plans{}).Where("id = ?", pkID).Updates(&plan)
	rowsEffected := gorm.RowsAffected
	log.Infof("rowsEffected : %v ", rowsEffected)
	return rowsEffected, nil
}

func DeletePlan(ctx context.Context, pkID int) (int64, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("err in dbConnection :****")
		return 0, err
	}
	var plan models.Plans
	result := dbConnection.Where("id = ?", pkID).Delete(&plan)
	if result.Error != nil {
		log.Info("***********failed to delete plan")
		return 0, result.Error
	}
	rowsEff := result.RowsAffected
	log.Infof("rowsEff : %v ", rowsEff)
	return rowsEff, nil
}
