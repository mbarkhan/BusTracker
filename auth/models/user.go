package models

import (
	"gorm.io/gorm"
)

type UserType struct {
	ID       uint8  `gorm:"primarykey"`
	UserType string `json:"user_type"`
	Users    []User `gorm:foreignKey:FUserType`
}

type User struct {
	gorm.Model
	// Auto ID added by gorm
	Name         string `json:"name"`
	Family       string `json:"family"`
	NationalId   string `json:"nationalID" gorm:"unique"` //This field should be unique
	Mobile       string `json:"mobile" gorm:"unique"`     //This field should be unique
	Password     string `json:"password"`
	FCompanyId   uint   `json:"company_id"`
	FUserTypeId  uint8  `json:"user_type"` //Initialize user role  0: public_paasenger  1: Admin //100: is reserved for Driver in Vehicles
	ProfileImage string `json:"profile_image"`
	IdCardImage  string `json:"id_image"`

	UserCredits      []UserCredit  `gorm:foreignKey:FUserId`
	Vehicles         []Vehicle     `gorm:foreignKey:FUserId`
	UserVehicles     []UserVehicle `gorm:foreignKey:FUserId`
	RegadminVehicles []UserVehicle `gorm:foreignKey:FRegadminId`
}
