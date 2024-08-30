package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type Account struct {
	owner   string
	balance int
}

var NoMoney = errors.New("Can't withdraw")

// create struct
func NewCreate(owner1 string) *Account {

	account := Account{owner: owner1, balance: 0}
	return &account
}

func (a *Account) Deposit(amount int) {
	fmt.Println("Goona deposit", amount)
	a.balance += amount

}

func (a *Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return NoMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {

	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner

}

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas:", a.Balance())
}
