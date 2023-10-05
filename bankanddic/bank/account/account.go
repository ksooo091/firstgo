package account

import (
	"errors"
	"fmt"
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

// ChangeName of Account
func (a *Account) ChangeName(newName string) {
	a.name = newName
}

// ShowName of Account
func (a Account) ShowName() string {
	return a.name
}

// 이름이 String 이여야함
// String about when call Account
func (a Account) String() string {
	return fmt.Sprint(a.ShowName(), "의 계좌 \n잔액 : ", a.ShowBalance())
}
