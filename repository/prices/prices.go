package prices

import (
	"context"

	bcontext "github.com/bubblepays/context"
	"github.com/bubblepays/database"
	"github.com/bubblepays/log"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)

func GetPrices(ctx context.Context) ([]models.Price, error) {
	//get dbconnection from contex
	dbConnection, err := bcontext.GetDBConnection(ctx)
	if err != nil {
		log.Errorft(ctx, "error in db connection %v ", err)
		return nil, err
	}
	var prices []models.Price
	if err := dbConnection.Find(&prices).Error; err != nil {
		log.Infof("could not get prices data %v ", err)
		return nil, err
	}
	log.Infof("prices : %v ", &prices)
	return prices, nil
}

func GetPricesByID(ctx context.Context, pkID int) (*models.Price, error) {
	log.Infof("This is GetPricesByID Repo")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("err %v ", err)
		return nil, err
	}
	var price models.Price
	if err := dbConnection.Where("id = ?", pkID).First(&price).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Info("plan not found")
		} else {
			log.Infof("Error: %v\n", err)
		}
		return nil, err
	}
	log.Infof("price : %v ", &price)
	return &price, nil
}

func SavePrice(ctx context.Context, price models.Price) (*models.Price, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("error in db connecton  ")
		return nil, err
	}
	ormDb := dbConnection.Create(&price)
	if ormDb.Error != nil {
		log.Infof("Error in inserting data in price %v ", ormDb.Error)
		return nil, ormDb.Error
	}
	rowsEffected := ormDb.RowsAffected
	log.Infof("Number of rows effected are: %v ", rowsEffected)
	return &price, nil
}

func UpdatePrice(ctx context.Context, pkID int, prices *models.Price) (*models.Price, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("error in db connecton  ")
		return nil, err
	}
	gorm := dbConnection.Model(&models.Price{}).Where("id = ?", pkID).Updates(&prices)
	if gorm != nil && gorm.Error != nil {
		log.Infof("error in updating price data")
		return nil, err
	}
	gorm.Find(&prices)
	return prices, nil
}

func DeletePrice(ctx context.Context, pkID int) (int64, error) {
	log.Infof("This is DeletePrice Repo")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Infof("error in db connecton  ")
		return 0, err
	}
	var price models.Price
	ormDb := dbConnection.Where("id=?", pkID).Delete(&price)
	rowsEffected := ormDb.RowsAffected
	return rowsEffected, nil
}
