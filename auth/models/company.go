package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	//Auto ID
	Name        string       `json:"name"`
	Longitude   float64      `json:"longitude"`
	Latitude    float64      `json:"latitude"`
	Address     string       `json:"address"`
	Description string       `json:"description"`
	FCityId     int          `json:"city_id"`
	Users       []User       `gorm:foreignKey:FCompanyId`
	UserCredits []UserCredit `gorm:foreignKey:FCompanyId`
}
