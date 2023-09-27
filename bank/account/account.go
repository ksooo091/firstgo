package account

import (
	"errors"
)

// Account Struct
type Account struct {
	// 대문자로 하면 값 Public 이 됨
	name    string
	balance int
}

var errNoMoney = errors.New("no money")

// NewAccount creates Account
func NewAccount(name string) *Account {
	account := Account{name: name, balance: 0}
	return &account
}

// Deposit x amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// ShowBalance about account
func (a Account) ShowBalance() int {
	return a.balance
}

// Withdraw x amount from account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
