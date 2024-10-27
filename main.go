package main

import (
	"github.com/go-resty/resty/v2"
	"github.com/ricardofabila/arithmetic-calculator-backend/controllers"
	"github.com/ricardofabila/arithmetic-calculator-backend/database"
	"github.com/ricardofabila/arithmetic-calculator-backend/routes"
	"github.com/ricardofabila/arithmetic-calculator-backend/services"
	"log"
)

func main() {
	log.Println("Starting Arithmetic Calculator Backend...")

	log.Println("Connecting to the database...")
	database.ConnectDatabase("calculator.db")
	log.Println("Database connection established successfully.")

	// Seed operations
	database.SeedOperations(database.DB)
	log.Println("Seeded operations successfully.")

	// Initialize the Resty client
	restyClient := resty.New()

	// Create an instance of RealRandomStringService with the Resty client
	randomStringService := &services.RealRandomStringService{
		Client: restyClient,
	}

	// Create an instance of the OperationController with the real RandomStringService
	operationController := &controllers.OperationController{
		RandomStringService: randomStringService,
	}

	// Set up the router
	log.Println("Setting up router...")
	r := routes.SetupRouter(operationController)
	log.Println("Router setup completed.")

	// Start the server and listen on port
	log.Println("Starting server on port 8080...")
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
