package controllers_test

import (
	"bytes"
	"github.com/ricardofabila/arithmetic-calculator-backend/services"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/ricardofabila/arithmetic-calculator-backend/controllers"
	"github.com/ricardofabila/arithmetic-calculator-backend/database"
	"github.com/ricardofabila/arithmetic-calculator-backend/models"
)

func setupTestDatabase() {
	// Connect to an in-memory SQLite database for testing
	database.ConnectDatabase(":memory:")

	// Create a test user
	testUser := models.User{
		Username: "testuser@example.com",
		Password: "password123",
		Balance:  100.0,
		Status:   "active",
	}
	database.DB.Create(&testUser)

	// Seed operations using the new reusable function
	database.SeedOperations(database.DB)
}

func TestPerformOperation_Success_WithMockRandomString(t *testing.T) {
	setupTestDatabase()

	mockRandomStringService := &services.MockRandomStringService{}

	// Create the controller instance with the mock service
	operationController := controllers.OperationController{
		RandomStringService: mockRandomStringService,
	}

	router := gin.Default()
	router.POST("/operation", func(c *gin.Context) {
		c.Set("user_id", uint(1)) // Mock user authentication
		operationController.PerformOperation(c)
	})

	// Create a test request for the "random_string" operation
	jsonBody := `{"operation": "random_string", "length": 10}`
	req, _ := http.NewRequest("POST", "/operation", bytes.NewBuffer([]byte(jsonBody)))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}

	expectedResult := `{"result":"MOCKSTRING"}`
	if w.Body.String() != expectedResult {
		t.Errorf("expected response %s, but got %s", expectedResult, w.Body.String())
	}
}

func TestGetRecords_Success(t *testing.T) {
	setupTestDatabase()

	router := gin.Default()
	router.GET("/records", func(c *gin.Context) {
		c.Set("user_id", uint(1)) // Mock user authentication
		controllers.GetRecords(c)
	})

	req, _ := http.NewRequest("GET", "/records?page=1&limit=10", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}

	if w.Body.Len() == 0 {
		t.Errorf("expected records, but got none")
	}
}

func TestDeleteRecord_Success(t *testing.T) {
	setupTestDatabase()

	// Create a record for testing deletion
	testRecord := models.Record{
		UserID:      1,
		OperationID: 1,
		Amount:      1.0,
		UserBalance: 99.0,
	}
	database.DB.Create(&testRecord)

	router := gin.Default()
	router.DELETE("/records/:id", func(c *gin.Context) {
		c.Set("user_id", uint(1))
		controllers.DeleteRecord(c)
	})

	req, _ := http.NewRequest("DELETE", "/records/1", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}

	// Verify the record has been deleted
	var record models.Record
	if err := database.DB.Where("id = ?", 1).First(&record).Error; err == nil {
		t.Errorf("expected record to be deleted, but found it")
	}
}
