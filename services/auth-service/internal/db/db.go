package db

import (
	"fmt"
	"log"

	"github.com/annguyen0511/Task-Management/services/auth-service/config"
	"github.com/annguyen0511/Task-Management/services/auth-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := db.AutoMigrate(&models.User{}); err != nil {
		return nil, err
	}
	log.Println("Connected to Postgres database")
	return db, nil

}

func Close(db *gorm.DB) {
	dbConn, err := db.DB()
	if err != nil {
		log.Println(err)
	}
	dbConn.Close()
	log.Println("Closed connection to Postgres database")
}
