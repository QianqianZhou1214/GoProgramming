package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int // Go allows you to create new types from existing ones.

type Stringer interface { // interface defined in fmt package
	String() string
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b) // let you define how your type is printed.
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in test is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance // can be (*w).balance, but Go permits us to write w.balance.
	// It's called struct pointer.
}

var ErrInsufficientFunds = errors.New("can't withdraw, insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
