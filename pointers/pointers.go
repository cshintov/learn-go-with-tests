package pointers

import "fmt"
import "errors"

// This is an error value. Don't use it. Instead use an error type.
// https://dave.cheney.net/2016/04/27/dont-just-check-errors-handle-them-gracefully
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

type Bitcoin int

func (b Bitcoin) String() string {
    return fmt.Sprintf("%d BTC", b)
}

// Implements Wallet
type Wallet struct {
    balance Bitcoin
}

func (w *Wallet) Balance() Bitcoin {
    return w.balance
}

func (w *Wallet) Deposit(deposit Bitcoin) {
    w.balance += deposit
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
    if amount > w.balance {
        return ErrInsufficientFunds
    }

    w.balance -= amount
    return nil
}
