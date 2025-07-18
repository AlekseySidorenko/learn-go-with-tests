package pointersanderrors

import "fmt"

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Printf("address of balance in Deposit is %p \n", &w.balance)
	return w.balance
}
