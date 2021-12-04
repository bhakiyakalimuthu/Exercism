package account

import (
	"sync"
)

var (
	mutex = sync.Mutex{}
)

// Define the Account type here.
type Account struct {
	balance int64
	closed  bool
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		balance: amount,
	}
}

func (a *Account) Balance() (int64, bool) {
	if !a.closed {
		return a.balance, true
	}
	return 0, false
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	if !a.closed {
		bal := a.balance + amount
		if bal >= 0 {
			a.balance = bal
			return a.balance, true
		}
	}
	return 0, false
}

func (a *Account) Close() (int64, bool) {
	mutex.Lock()
	defer mutex.Unlock()
	if a.closed {
		return 0, false
	}
	a.closed = true
	return a.balance, true
}
