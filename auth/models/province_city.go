package models

type  Province struct{
	ID uint `gorm:"primarykey"`
	Name string `json:"name"`
	Cities []City `gorm:"foreignkey:FProvinceId"`
}

type City struct{
	ID uint  `gorm:"primarykey"`
	FProvinceId uint `json:"f_province_id"`
	Name string `json:"name"`
	Companies []Company `gorm:"foreignkey:FCityId"`
}