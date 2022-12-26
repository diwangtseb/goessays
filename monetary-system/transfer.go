package main

import "context"

type AccountTransfer interface {
	AdjustAccountAmount(ctx context.Context, aid string, amount float64) error
}
