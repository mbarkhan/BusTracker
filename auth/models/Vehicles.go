package models

import "gorm.io/gorm"

type Vehicle struct {
	ID uint `gorm:"primarykey"`
	FUserID   uint    `json:"f_user_id"`
	VehicleName  string `json:"vehicle_name"`
	Color        string `json:"color"`
	VID          string `json:"v_id"` // Plak
	VehicleImage string `json:"vehicle_image"`

	VehicleLogs []VehicleLog `gorm:"foreignkey:FVehicleId"`
}


type VehicleLog struct{
	gorm.model
	FVehicleId uint `json:"f_vehicle_id"`
	DateTime time.Time `json:"date_time"`
	Longitude float64 `json:"Longitude"`
	Latitude float64 `json:"Latitude"`
	
}