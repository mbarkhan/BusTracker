package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model

	Name        string   `json:"name"`
	Location    Location `json:"location"`
	Description string   `json:"description"`
}
