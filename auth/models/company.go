package models

import (
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	//Auto ID
	Name        string  `json:"name"`
	AdminName   string  `json:"admin_name"`
	AdminFamily string  `jaon:"admin_family"`
	Mobile      string  `json:"mobile" gorm:"unique"`
	Longitude   float64 `json:"longitude"`
	Latitude    float64 `json:"latitude"`
	Description string  `json:"description"`
}
