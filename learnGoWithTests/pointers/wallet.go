package pointers

import "fmt"

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(amount int) {
	fmt.Printf("address of balance in test is %p \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance // can be (*w).balance, but Go permits us to write w.balance.
	// It's called struct pointer.
}
