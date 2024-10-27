package database

import (
	"github.com/ricardofabila/arithmetic-calculator-backend/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

// ConnectDatabase initializes a database connection
// (can be in-memory for testing or a file path)
func ConnectDatabase(dsn string) {
	database, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Automatically migrate models (create tables if they don't exist)
	database.AutoMigrate(&models.User{}, &models.Operation{}, &models.Record{})
	DB = database
}

// SeedOperations make sure to initialize the database with the operations if not present
func SeedOperations(db *gorm.DB) {
	operations := []models.Operation{
		{Type: "addition", Cost: 1.0},
		{Type: "subtraction", Cost: 1.0},
		{Type: "multiplication", Cost: 1.5},
		{Type: "division", Cost: 2.0},
		{Type: "square_root", Cost: 2.5},
		{Type: "random_string", Cost: 2.5},
	}

	for _, op := range operations {
		if err := db.Where("type = ?", op.Type).FirstOrCreate(&op).Error; err != nil {
			panic(err)
		}
	}
}
