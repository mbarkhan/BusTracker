package models

import (
	"gorm.io/gorm"
)

type Permission struct {
	gorm.Model

	Name        string `json:"name"`
	Route       string `json:"route"`
	Description string `json:"description"`
}
