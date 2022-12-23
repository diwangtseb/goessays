package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	moneTrayA := NewMoneTrayActor("moneTrayA")
	for i := 0; i < 1000; i++ {
		go func() {
			moneTrayA.TryTrans(context.TODO(), "a", "b", 100)
			moneTrayA.RollTrans(context.TODO(), "b", "a", 100)
			moneTrayA.TryTrans(context.TODO(), "a", "b", 100)
			moneTrayA.RollTrans(context.TODO(), "b", "a", 100)
		}()
	}
	time.Sleep(time.Second * 5)
	go func() {
		r := moneTrayA.LookupAccountAmount(context.TODO())
		for _, v := range r {
			fmt.Println("account:", v.name, v.amount)
		}
	}()
	time.Sleep(time.Second * 5)
}

type MoneTaryTransfer interface {
	TryTrans(ctx context.Context, fromAccount string, toAcccount string, amount float64) error
	RollTrans(ctx context.Context, fromAccount string, toAcccount string, ammount float64) error
}

type Account struct {
	name   string
	amount float64
}

type MoneTrayActor struct {
	bankAccount    *Account
	allUserAccount *Account
	lock           sync.RWMutex
}

func getAllUserAccountByBankAccountName(bankAccountName string) string {
	userBankName := fmt.Sprintf("%s:user:account", bankAccountName)
	return userBankName
}

func NewMoneTrayActor(bankAccountName string) *MoneTrayActor {
	return &MoneTrayActor{
		bankAccount: &Account{
			name: bankAccountName,
		},
		allUserAccount: &Account{
			name: getAllUserAccountByBankAccountName(bankAccountName),
		},
	}
}

// RollTrans implements MoneTaryTransfer
func (mta *MoneTrayActor) RollTrans(ctx context.Context, fromAccount string, toAcccount string, ammount float64) error {
	if toAcccount != mta.bankAccount.name {
		toAcccount = mta.bankAccount.name
	}
	fmt.Println(fromAccount, "-->", toAcccount, ammount)
	mta.lock.Lock()
	defer mta.lock.Unlock()
	mta.bankAccount.amount -= ammount
	mta.bankAccount.amount += ammount
	return nil
}

// TryTrans implements MoneTaryTransfer
func (mta *MoneTrayActor) TryTrans(ctx context.Context, fromAccount string, toAcccount string, amount float64) error {
	if fromAccount != mta.bankAccount.name {
		fromAccount = mta.bankAccount.name
	}
	fmt.Println(fromAccount, "-->", toAcccount, amount)
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
	as = append(as, &(*mta.allUserAccount))
	as = append(as, &(*mta.bankAccount))
	return as
}
