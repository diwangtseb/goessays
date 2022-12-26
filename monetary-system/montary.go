package main

import (
	"context"
	"fmt"
	"sync"
)

type MoneTaryTransfer interface {
	TryTrans(ctx context.Context, fromAccountId string, toAcccountId string, amount float64) error
	RollTrans(ctx context.Context, fromAccountId string, toAcccountId string, ammount float64) error
}

type MoneTrayActor struct {
	bankAccount    *Account
	allUserAccount *Account
	lock           sync.RWMutex
}

func getBatchUserAccountByBankAccountId(bankAccountId string) string {
	batchUserBankId := fmt.Sprintf("%s:batch:user:account", bankAccountId)
	return batchUserBankId
}

func NewMoneTrayActor(bankAccount *Account) *MoneTrayActor {
	return &MoneTrayActor{
		bankAccount: bankAccount,
		allUserAccount: &Account{
			id: getBatchUserAccountByBankAccountId(bankAccount.id),
		},
	}
}

// RollTrans implements MoneTaryTransfer
func (mta *MoneTrayActor) RollTrans(ctx context.Context, fromAccountId string, toAcccountId string, ammount float64) error {
	if !mta.bankAccount.isValidAccount(toAcccountId) {
		toAcccountId = mta.bankAccount.id
	}
	fmt.Println(fromAccountId, "-->", toAcccountId, ammount)
	mta.lock.Lock()
	defer mta.lock.Unlock()
	mta.bankAccount.amount -= ammount
	mta.bankAccount.amount += ammount
	return nil
}

// TryTrans implements MoneTaryTransfer
func (mta *MoneTrayActor) TryTrans(ctx context.Context, fromAccountId string, toAcccountId string, amount float64) error {
	if !mta.bankAccount.isValidAccount(fromAccountId) {
		fromAccountId = mta.bankAccount.id
	}
	fmt.Println(fromAccountId, "-->", toAcccountId, amount)
	mta.lock.Lock()
	defer mta.lock.Unlock()
	mta.bankAccount.amount += amount
	mta.allUserAccount.amount -= amount
	return nil
}

func (mta *MoneTrayActor) LookupAccountAmount(ctx context.Context) []*Account {
	mta.lock.RLock()
	defer mta.lock.RUnlock()
	as := make([]*Account, 0)
	allUserAccount := mta.allUserAccount
	bankAccount := mta.bankAccount
	as = append(as, allUserAccount)
	as = append(as, bankAccount)
	return as
}
