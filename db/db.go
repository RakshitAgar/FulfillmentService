package db

import (
	"fullfilmentService/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	var err error
	dsn := "host=localhost user=delivery_user password=your_password dbname=fulfillmentService port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	// Auto-migrate the Delivery model
	err = DB.AutoMigrate(&model.Delivery{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v\n", err)
	}
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatalf("Failed to get DB instance: %v\n", err)
	}
	sqlDB.Close()
}
