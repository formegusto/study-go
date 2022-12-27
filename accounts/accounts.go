package accounts

import (
	"fmt"

	"github.com/formegusto/study-go/myerrors"
)

// Account struct
type account struct {
	owner string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *account {
	a := account{owner: owner, balance: 0}

	return &a
}

// Deposit x amount on your account
func (a *account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a account) Balance() int {
	return a.balance
}

// 출금 : Withdraw x amount from your account
func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return myerrors.NoMoney
	}
	a.balance -= amount
	return nil
}

// Change Owner of the account
func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a account) Owner() string {
	return a.owner
}

/*
Python의 __str__과 같은 역할
*/
func (a account) String() string {
	return fmt.Sprint(a.Owner(),"'s account. \nHas: $", a.Balance())
}