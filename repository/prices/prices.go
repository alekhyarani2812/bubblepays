package prices

import (
	"fmt"
	"log"

	"github.com/bubblepays/database"
	"github.com/bubblepays/models"
	"gorm.io/gorm"
)

func GetPrices() ([]models.Price, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err)
		return nil, err
	}
	var prices []models.Price
	if err := dbConnection.Find(&prices).Error; err != nil {
		log.Println("could not get prices data", err)
		return nil, err
	}
	log.Println("prices : ", &prices)
	return prices, nil
}

func GetPricesById(pkId int) (*models.Price, error) {
	fmt.Println("This is GetPricesById Repo")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		log.Printf("err %v ", err)
		return nil, err
	}
	var price models.Price
	if err := dbConnection.Where("id = ?", pkId).First(&price).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("plan not found")
		} else {
			log.Printf("Error: %v\n", err)
		}
		return nil, err
	}
	log.Println("price : ", &price)
	return &price, nil
}

func SavePrice(price models.Price) (*models.Price, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connecton  ")
		return nil, err
	}
	ormDb := dbConnection.Create(&price)
	if ormDb.Error != nil {
		fmt.Println("Error in inserting data in price", ormDb.Error)
		return nil, ormDb.Error
	}
	rowsEffected := ormDb.RowsAffected
	fmt.Printf(" %d rows effected : ", rowsEffected)
	return &price, nil
}

func UpdatePrice(pkId int, prices *models.Price) (*models.Price, error) {
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connecton  ")
		return nil, err
	}
	gorm := dbConnection.Model(&models.Price{}).Where("id = ?", pkId).Updates(&prices)
	if gorm != nil && gorm.Error != nil {
		fmt.Println("error in updating price data")
		return nil, err
	}
	gorm.Find(&prices)
	return prices, nil
}

func DeletePrice(pkId int) (int64, error) {
	fmt.Println("This is DeletePrice Repo")
	dbConnection, err := database.GetORMDbConnection()
	if err != nil {
		fmt.Println("error in db connecton  ")
		return 0, err
	}
	var price models.Price
	ormDb := dbConnection.Where("id=?", pkId).Delete(&price)
	rowsEffected := ormDb.RowsAffected
	return rowsEffected, nil
}
