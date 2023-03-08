package models

import (
	"time"

	"gorm.io/gorm"
)

type Company_Users struct {
	gorm.Model
	// Auto ID
	FCompanyID int       `json:"f_company_id"`
	FUserID    int       `json:"f_user_id"`
	ConfirmAt  time.Time `json:"confirm_at"`
	Longitude  float64   `json:"longitude"` // The Location of user to get on or off
	Latitude   float64   `json:"latitude"`
}
