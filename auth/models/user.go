package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	// Auto ID added by gorm
	FCompanyId uint `json:"f_company_id"`
	FUserTypeId uint `json:"f_user_type_id"`
	Name       string `json:"name"`
	Family     string `json:"family"`
	NationalID string `json:"nationalID" gorm:"unique"` //This field should be unique
	Password   string `json:"password"`
	Mobile     string `json:"mobile" gorm:"unique"` //This field should be unique
	PrifileImage string `json:"profile_image"`
	IdCardImage string `json:"idcard_image"`

	Vehicles []Vehicle `gorm:"foreignkey:FUserID"`
	UserCredits []UserCredit `gorm:"foreignkey:FUserID"`
}

type UserType struc{
	ID uint `gorm:"primarykey"`
	UserType string `json:"user_type"`
	Users []user `gorm:"foreignkey:FUserTypeId"`
}


