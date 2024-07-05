package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	//This is a blank import. It is a mysql driver library
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GetDbConnection get DBConnection
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

// GetORMDbConnection get orm db connection
func GetORMDbConnection() (*gorm.DB, error) {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,       // Disable color
		},
	)
	//echo $DB_CONNECTION
	//dsn := "root:root@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := os.Getenv("DB_CONNECTION")
	fmt.Printf("Mysql DSN DB Connection %v ", dsn)
	//dsn := "root:PASSWORD@tcp(192.168.49.2:3306)/todo"
	//dsn := "root:PASSWORD@tcp(192.168.49.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"

	//"root:PASSWORD@tcp(mysql-local:3306)/todo"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		fmt.Println("database connection error")
		return nil, err
	}
	// Enable debug mode
	db = db.Debug()
	return db, nil
}
