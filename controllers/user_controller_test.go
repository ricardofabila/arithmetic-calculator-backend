package controllers

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/ricardofabila/arithmetic-calculator-backend/database"
	"github.com/ricardofabila/arithmetic-calculator-backend/models"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// setupTestDatabase initializes an in-memory database and seeds data.
func setupTestDatabase() {
	// Initialize the in-memory database
	database.ConnectDatabase(":memory:")

	// Seed initial data for the tests
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	database.DB.Create(&models.User{
		Username: "testuser@example.com",
		Password: string(hashedPassword),
		Status:   "active",
		Balance:  100.0,
	})
}

// setupRouter sets up a Gin router with the specified routes.
func setupRouter(routeSetup func(r *gin.Engine)) *gin.Engine {
	router := gin.Default()
	routeSetup(router)
	return router
}

// performRequest performs an HTTP request on the provided router.
func performRequest(router *gin.Engine, method, path string, body []byte) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}

func TestRegisterUser(t *testing.T) {
	setupTestDatabase()

	// Set up a new router with the /register route
	router := setupRouter(func(r *gin.Engine) {
		r.POST("/register", RegisterUser)
	})

	// Test a successful user registration
	jsonBody := `{"username": "newuser@example.com", "password": "password123"}`
	w := performRequest(router, "POST", "/register", []byte(jsonBody))

	// Check if the response is OK (200)
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}

	// Verify user was added to the in-memory database
	var user models.User
	if err := database.DB.Where("username = ?", "newuser@example.com").First(&user).Error; err != nil {
		t.Errorf("expected user to be created, but got error: %v", err)
	}
}

func TestLoginUser(t *testing.T) {
	setupTestDatabase()

	// Set up a new router with the /login route
	router := setupRouter(func(r *gin.Engine) {
		r.POST("/login", LoginUser)
	})

	// Test a successful login request
	jsonBody := `{"username": "testuser@example.com", "password": "password123"}`
	w := performRequest(router, "POST", "/login", []byte(jsonBody))

	// Check if the response status is OK (200)
	if w.Code != http.StatusOK {
		t.Errorf("expected status OK, got %v", w.Code)
	}

	// Test a failed login request with incorrect credentials
	wrongJsonBody := `{"username": "testuser@example.com", "password": "wrongpassword"}`
	wrongW := performRequest(router, "POST", "/login", []byte(wrongJsonBody))

	// Check if the response status is Unauthorized (401)
	if wrongW.Code != http.StatusUnauthorized {
		t.Errorf("expected status Unauthorized, got %v", wrongW.Code)
	}
}
