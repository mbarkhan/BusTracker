package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Auto ID added by gorm
	Name       string `json:"name"`
	Family     string `json:"family"`
	NationalID string `json:"nationalID" gorm:"unique"` //This field should be unique
	Password   string `json:"password"`
	Mobile     string `json:"mobile" gorm:"unique"` //This field should be unique
}
