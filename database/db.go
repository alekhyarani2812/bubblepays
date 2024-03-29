package database

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// get DBConnection
func GetDbConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/todo")
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error db connection : ", err.Error())
	} else {
		fmt.Println("Success db connection ")
	}
	return db
}

// get orm db connection
func GetORMDbConnection() (*gorm.DB, error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		//panic("failed to connect to database")
		log.Printf("database connection error")
		return nil, err
	}
	fmt.Println(db)
	return db, nil
}
