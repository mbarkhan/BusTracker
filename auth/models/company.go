package models

type Company struct {
	ID          uint         `gorm:"primarykey"`
	FCityId     uint         `json:"f_city_id"`
	Name        string       `json:"name"`
	Longitude   float64      `json:"longitude"`
	Latitude    float64      `json:"latitude"`
	Address     string       `json:"address"`
	Description string       `json:"description"`
	Users       []User       `gorm:"foreignkey:FCompanyId"`
	UserCredits []UserCredit `gorm:"foreignkey:FCompanyId"`
}
