package models

import (
	"gorm.io/gorm"
)

type Tariff struct {
	ID                uint16       `gorm:"primarykey"`
	TariffDescription string       `json:"tariff_description"`
	Days              uint16       `json:"days"`
	Price             uint32       `json:"price"`
	Active            bool         `json:"actve"`
	UserCredits       []UserCredit `gorm:foreignKey:FTariffId`
}

type UserCredit struct {
	gorm.Model
	// Auto ID added by gorm
	//DepositDate time.Time `json:"depositDate"` ّIS Embeded By gorm.Model
	FUserID    uint    `json:"f_user_id"`
	FCompanyId uint    `json:"f_company_id"`
	FTariffId  uint    `json:"f_tariff_id"`
	Deposit    float32 `json:"deposit"` //The Amount of money that has been deposited
	BankTx     string  `json:"bank_Tx"` //Banking Transaction ID

}
