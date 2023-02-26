package models

import "gorm.io/gorm"

type Vehicles struct {
	gorm.Model
	//Auto ID
	FCompanyID   int    `json:"f_company_id"`
	Vehiclename  string `json:"vehicle_name"`
	Color        string `json:"color"`
	VID          string `json:"v_id"`
	PathName     string `json:"Path_Name"`
	DriverName   string `json:"driver_name"`
	DriverFamily string `json:"driver_family"`
	DriverMobile string `json:"driver_mobile"`
	Password     string `json:"password"`
	DriverImage  string `json:"driver_image"`
	VehicleImage string `json:"vehicle_image"`
}
