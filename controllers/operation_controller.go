package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/ricardofabila/arithmetic-calculator-backend/database"
	"github.com/ricardofabila/arithmetic-calculator-backend/models"
	"github.com/ricardofabila/arithmetic-calculator-backend/services"
	"net/http"
	"strconv"
	"time"
)

type OperationRequest struct {
	Operation string   `json:"operation" binding:"required"`
	Number1   *float64 `json:"number1"`
	Number2   *float64 `json:"number2"`
	Length    *int     `json:"length"` // Length for the random string
}

type OperationController struct {
	RandomStringService services.RandomStringService
}

func (oc *OperationController) PerformOperation(c *gin.Context) {
	var req OperationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Get the user ID from the request context (set by the JWT middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var user models.User
	if err := database.DB.First(&user, userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	var operation models.Operation
	if err := database.DB.Where("type = ?", req.Operation).First(&operation).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid operation type"})
		return
	}

	// Check if the user has sufficient balance for the operation
	if user.Balance < operation.Cost {
		c.JSON(http.StatusPaymentRequired, gin.H{"error": "Insufficient balance"})
		return
	}

	var result string
	var err error
	switch req.Operation {
	case "addition", "subtraction", "multiplication", "division":
		if req.Number1 == nil || req.Number2 == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Both number1 and number2 are required for this operation"})
			return
		}
		result, err = services.PerformArithmeticOperation(req.Operation, *req.Number1, *req.Number2)
	case "square_root":
		if req.Number1 == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "number1 is required for square root operation"})
			return
		}
		result, err = services.Sqrt(*req.Number1)
	case "random_string":
		length := 10 // Default length
		if req.Length != nil {
			length = *req.Length
		}
		result, err = oc.RandomStringService.GetRandomString(length)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported operation"})
		return
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Deduct the cost from the user's balance
	user.Balance -= operation.Cost
	if err := database.DB.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user balance"})
		return
	}

	record := models.Record{
		OperationID:     operation.ID,
		UserID:          user.ID,
		Amount:          operation.Cost,
		UserBalance:     user.Balance,
		OperationResult: result,
		Date:            time.Now().Format(time.RFC3339),
	}

	if err := database.DB.Create(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": result})
}

func GetRecords(c *gin.Context) {
	// Get the user ID from the request context (set by the JWT middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	search := c.Query("search")

	// Start building the query
	query := database.DB.Where("user_id = ?", userID)

	// Apply search filter if search parameter is present
	if search != "" {
		query = query.Where("operation_result LIKE ?", "%"+search+"%")
	}

	// Get the total count of records for the current user
	var totalCount int64
	if err := query.Model(&models.Record{}).Count(&totalCount).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch total record count"})
		return
	}
	totalPages := (totalCount + int64(limit) - 1) / int64(limit)

	var records []models.Record
	// using offset pagination and not a cursor to keep things simple
	if err := query.Preload("Operation").Limit(limit).Offset(offset).Find(&records).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch records"})
		return
	}

	responseRecords := []map[string]interface{}{}
	for _, record := range records {
		responseRecords = append(responseRecords, map[string]interface{}{
			"id":        record.ID,
			"amount":    record.Amount,
			"date":      record.Date,
			"result":    record.OperationResult,
			"operation": record.Operation.Type,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"records":    responseRecords,
		"totalPages": totalPages,
	})
}

func DeleteRecord(c *gin.Context) {
	// Get the user ID from the request context (set by the JWT middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// Get the record ID from the URL parameter
	recordID := c.Param("id")

	var record models.Record
	if err := database.DB.Where("id = ? AND user_id = ?", recordID, userID).First(&record).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Record not found"})
		return
	}

	if err := database.DB.Delete(&record).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete record"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Record deleted successfully"})
}
