package models

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func init() {

	err := godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbDriver := os.Getenv("DB_DRIVER")

	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)

	fmt.Printf("Connecting to %s database", dbDriver)
	conn, err := gorm.Open(dbDriver, dbURI)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", dbDriver)
		log.Fatal("Error:", err)
	}

	db = conn
	db.Debug().AutoMigrate(&User{}, &Company{})
	db.Model(&User{}).AddForeignKey("company_id", "companies(id)", "CASCADE", "CASCADE")
}

func GetDB() *gorm.DB {
	return db
}
