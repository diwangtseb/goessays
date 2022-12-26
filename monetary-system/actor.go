package main

import (
	"context"
	"errors"
)

type actor struct {
	fromAT AccountTransfer
	toAT   AccountTransfer
}

// RollTrans implements MoneTaryTransfer
func (a *actor) RollTrans(ctx context.Context, fromAccountId string, toAcccountId string, amount float64) error {
	// start roll transfer
	delAmount := -amount
	err := a.fromAT.AdjustAccountAmount(ctx, fromAccountId, delAmount)
	if err != nil {
		return errors.New("adjust account amount error")
	}
	incrAmount := amount
	err = a.toAT.AdjustAccountAmount(ctx, toAcccountId, incrAmount)
	if err != nil {
		return errors.New("adjust account amount error")
	}
	return nil
}

// TryTrans implements MoneTaryTransfer
func (a *actor) TryTrans(ctx context.Context, fromAccountId string, toAcccountId string, amount float64) error {
	// start transfer
	delAmount := -amount
	err := a.fromAT.AdjustAccountAmount(ctx, fromAccountId, delAmount)
	if err != nil {
		return errors.New("adjust account amount error")
	}
	incrAmount := amount
	err = a.toAT.AdjustAccountAmount(ctx, toAcccountId, incrAmount)
	if err != nil {
		return errors.New("adjust account amount error")
	}
	return nil
}

func NewMoneTaryTransferActor(formAccountTransfer AccountTransfer, toAccountTransfer AccountTransfer) MoneTaryTransfer {
	return &actor{
		fromAT: formAccountTransfer,
		toAT:   toAccountTransfer,
	}
}

var _ MoneTaryTransfer = (*actor)(nil)
