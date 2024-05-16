package database

import (
	"database/sql"
	"fmt"
	"log"
	//This is a blank import. It is a mysql driver library
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//GetDbConnection get DBConnection
func GetDbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/todo")
	if err != nil {
		log.Fatal(err.Error())
		fmt.Printf("Error db connection : %v ", err.Error())
	} else {
		log.Printf("Success db connection ")
	}
	return db
}

//GetORMDbConnection get orm db connection
func GetORMDbConnection() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database connection error")
		return nil, err
	}
	return db, nil
}
