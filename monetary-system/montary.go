package main

import (
	"context"
)

type MoneTaryTransfer interface {
	TryTrans(ctx context.Context, fromAccountId string, toAcccountId string, amount float64) error
	RollTrans(ctx context.Context, fromAccountId string, toAcccountId string, ammount float64) error
}
