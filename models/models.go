package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `gorm:"unique;not null" json:"username"`
	Password string  `gorm:"not null" json:"password"`
	Status   string  `gorm:"default:active" json:"status"`
	Balance  float64 `gorm:"default:50" json:"balance"`
}

type Operation struct {
	ID   uint    `gorm:"primaryKey" json:"id"`
	Type string  `gorm:"unique;not null" json:"type"`
	Cost float64 `gorm:"not null" json:"cost"`
}

type Record struct {
	gorm.Model
	OperationID     uint      `json:"operationId"`
	UserID          uint      `json:"userId"`
	Amount          float64   `json:"amount"`
	UserBalance     float64   `json:"userBalance"`
	OperationResult string    `json:"operationResult"` // string since it can be a number or a string
	Date            string    `json:"date"`
	Operation       Operation `json:"operation" gorm:"foreignKey:OperationID"`
}
