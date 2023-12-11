package database

import (
	"log"
	"github.com/yeshwanthjampala/nearby_api/models"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connectionString := "host=" + dbHost + " port=" + dbPort + " user=" + dbUsername + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

	database, err := gorm.Open("postgres", connectionString)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	database.DB().SetMaxIdleConns(10)
	database.DB().SetMaxOpenConns(100)

	db = database

	// AutoMigrate any models if required
	db.AutoMigrate(&models.Location{}, &models.SearchParams{}, &models.TripCost{}, &models.UserLocation{})

	return db
}

func GetDB() *gorm.DB {
	return db
}

func CloseDB() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatal("Error closing database: ", err)
		}
	}
}
