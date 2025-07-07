package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return (*w).balance
}

// *Wallet is a struct pointer
func (w *Wallet) Deposit(amount Bitcoin) {
	// The dereferencing here is not necessary in Go
	// It happens automatically for struct pointers
	// May as well have been:
	// w.balance += amount
	(*w).balance += amount
}

// var keyword allows defintion of values global to the package
var ErrEasterEgg = errors.New("to the moon 🚀️🚀️🚀️")

func ErrInsufficientFundsOnWithdraw(amount, balance Bitcoin) error {
	return fmt.Errorf("insufficient balance – attempting to withdraw: %s; current balance: %s", amount, balance)
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	balance := w.Balance()

	if w.Balance() < amount {
		return ErrInsufficientFundsOnWithdraw(amount, balance)
	}

	(*w).balance -= amount

	return nil
}
