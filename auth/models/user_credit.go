package models

import (
	"time"

	"gorm.io/gorm"
)

type User_Credit struct {
	gorm.Model
	// Auto ID added by gorm
	UserID      int       `json:"userID"`
	CompanyID   int       `json:"companyID"`
	Deposit     float32   `json:"deposit"` //The Amount of money that has been deposited
	BankTx      string    `json:"bankTx"`  //Banking Transaction ID
	DepositDate time.Time `json:"depositDate"`
	CreditDays  int       `json:"creditDays"` //
}
