package main

import "context"

type UserAccountManager interface {
	CreateUserAccount(userAccount *Account) error
	GetAccountByUid(uid string) *Account
	AccountTransfer
}

type BankAccountManager interface {
	CreateBankAccount(bankAccount *Account) error
	GetAccountByBid(bid string) *Account
	AccountTransfer
}

var _ BankAccountManager = (*bankAccountManager)(nil)
var _ UserAccountManager = (*userAccountManager)(nil)

type Account struct {
	id     string
	amount float64
}

func (a *Account) deductmount(amount float64) {
	a.amount--
}

func (a *Account) isValidAccount(id string) bool {
	return a.id == id
}

type userAccountManager struct{}

func NewUserAccountManager() UserAccountManager {
	return &userAccountManager{}
}

// AdjustAccountAmount implements AccountTransfer
func (uam *userAccountManager) AdjustAccountAmount(ctx context.Context, aid string, amount float64) error {
	account := uam.GetAccountByUid(aid)
	var userAccount *Account
	if account == nil {
		//crate account
		uid := aid
		userAccount = &Account{
			id:     uid,
			amount: amount,
		}
		err := uam.CreateUserAccount(userAccount)
		if err != nil {
			return err
		}
	}
	account.deductmount(amount)
	return nil
}

// CreateUserAccount implements UserAccountManager
func (*userAccountManager) CreateUserAccount(userAccount *Account) error {
	return nil
}

// GetAccountByUid implements UserAccountManager
func (*userAccountManager) GetAccountByUid(uid string) *Account {
	return &Account{
		id: uid,
	}
}

type bankAccountManager struct{}

func NewBankAccountManager() BankAccountManager {
	return &bankAccountManager{}
}

// AdjustAccountAmount implements AccountTransfer
func (bam *bankAccountManager) AdjustAccountAmount(ctx context.Context, aid string, amount float64) error {
	account := bam.GetAccountByBid(aid)
	var userAccount *Account
	if account == nil {
		//crate account
		uid := aid
		userAccount = &Account{
			id:     uid,
			amount: amount,
		}
		err := bam.CreateBankAccount(userAccount)
		if err != nil {
			return err
		}
	}
	account.deductmount(amount)
	return nil
}

// CreateBankAccount implements BankAccountManager
func (*bankAccountManager) CreateBankAccount(bankAccount *Account) error {
	return nil
}

// GetAccountByBid implements BankAccountManager
func (*bankAccountManager) GetAccountByBid(bid string) *Account {
	return &Account{
		id:     bid,
		amount: 99999,
	}
}
