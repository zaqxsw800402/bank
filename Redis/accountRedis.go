package Redis

import (
	"fmt"
	"red/domain"
	"time"
)

type AccountResponse struct {
	AccountId   uint
	CustomerId  uint
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func GetAccount(id string) *AccountResponse {
	key := fmt.Sprintf("%s:%s:response", Account, id)
	res := RC.HGetAll(ctx, key)
	result, _ := res.Result()
	if _, err := result["AccountId"]; !err {
		return nil
	}

	var account AccountResponse
	if err := res.Scan(&account); err != nil {
		panic(err)
	}

	return &account
}

func SaveAccount(c *domain.Account) {
	key := fmt.Sprintf("%s:%d:response", Account, c.AccountId)
	RC.HSet(ctx, key,
		"AccountId", c.AccountId,
		"CustomerId", c.CustomerId,
		"OpeningDate", c.OpeningDate,
		"AccountType", c.AccountType,
		"Amount", c.Amount,
		"Status", c.Status,
	)
	RC.Expire(ctx, key, time.Minute)
}