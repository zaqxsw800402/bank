package dto

import (
	"red/cmd/api/errs"
	"strings"
)

type NewAccountRequest struct {
	CustomerId  uint    `json:"customer_id"`
	AccountType string  `json:"account_type" binding:"required"`
	Amount      float64 `json:"amount" binding:"required"`
}

type NewAccountResponse struct {
	AccountId uint `json:"account_id"`
}

// Validate 檢查body裡的資料是否有效
func (r NewAccountRequest) Validate() *errs.AppError {
	// 初始帳戶低於5000元，則判定無效
	if r.Amount < 5000 {
		return errs.NewValidationError("To open a new account you need to deposit at least 5000")
	}
	// 判斷是否有效
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewValidationError("Account type must be saving or checking")
	}
	return nil

}