package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Auto ID added by gorm
	Name   string `json:"name"`
	Family string `json:"family"`
	//Username       string    `json:"username"`
	NationalID string `json:"nationalID" gorm:"unique"` //This field should be unique
	//BirthDate      time.Time `json:"birthDate"`
	//Email          string    `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile" gorm:"unique"` //This field should be unique
	//Phone          string    `json:"phone"`
	//Shaba          string    `json:"sheba"`
	VerifiedAt time.Time `json:"verified"`
	Longitude  float64   `json:"longitude"`
	Latitude   float64   `json:"latitude"`
	//DetailedbookID int       `json:"detailedbookID"`

}
