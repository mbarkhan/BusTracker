package models

import "gorm.io/gorm"

type Pathes struct {
	gorm.Model
	//Auto ID
	FCompanyID  int    `json:"f_company_id"`
	PathName    string `json:"path_name"`
	Description string `json:"description"`
}
