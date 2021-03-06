package domain

import (
	"red/dto"
)

type Transaction struct {
	TransactionId   uint `gorm:"primaryKey;autoIncrement"`
	AccountId       uint
	Amount          float64
	NewBalance      float64
	TransactionType string
	TransactionDate string
}

func (t Transaction) IsWithdrawal() bool {
	if t.TransactionType == "withdrawal" || t.TransactionType == "transfer out" {
		return true
	}
	return false
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:     t.TransactionId,
		TransactionAmount: t.Amount,
		NewBalance:        t.NewBalance,
		TransactionType:   t.TransactionType,
		TransactionDate:   t.TransactionDate,
	}
}
