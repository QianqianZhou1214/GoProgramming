package pointers

import "fmt"

type Bitcoin int // Go allows you to create new types from existing ones.
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
