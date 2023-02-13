package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	Name        string  `json:"name"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Description string  `json:"description"`
}
