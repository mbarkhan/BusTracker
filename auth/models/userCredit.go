package models

import (
	"time"

	"gorm.io/gorm"
)

type UserCredit struct {
	gorm.Model
	// Auto ID added by gorm
	FUserID     uint      `json:"f_user_id"`
	FCompanyID  uint      `json:"f_company_id"`
	FTarrifID   uint      `json:"f_tariff_id"`
	BankTx      string    `json:"bankTx"` //Banking Transaction ID
	DepositDate time.Time `json:"depositDate"`
}

type Tariff struct {
	ID          uint         `gorm:"primarykey"`
	Tariff      string       `json:"tariff"`
	Days        uint16       `json:"days"`
	Active      bool         `json:"active"`
	UserCredits []UserCredit `gorm:"foreignkey:FTariffID"`
}
