package accounts

import (
	"errors"
	"fmt"
)

// Account Struct
type account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Cant withdraw")

// New Account creates
func NewAccount(owner string) *account{
	account := account{owner:owner,balance:0}
	return &account
}

//Deposit x amount on your account
func (a *account) Deposit(amount int){
	a.balance += amount
}

//Withdraw from account
func (a *account) Withdraw(amount int) error{
	if a.balance < amount{
		return errNoMoney
	}
	a.balance -= amount
	return nil
}
//Change owner
func (a *account) ChangeOwner(newOwner string){
	a.owner = newOwner
}

//Balance
func (a account) Balance() int{
	return a.balance
}

//Owner
func (a account) Owner() string{
	return a.owner
}

func (a account) String() string {
	return fmt.Sprint(a.Owner(),"'s account \nHas ",a.Balance())
}