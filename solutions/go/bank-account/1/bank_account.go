package bankaccount

import "sync"

type Account struct {
	amount int64
    isOpen bool
    mu sync.Mutex
}

func Open(amt int64) *Account {
    if amt < 0 {
        return nil
    }
    
	return &Account {
        amount: amt,
        isOpen: true,
    }
}

func (a *Account) Balance() (bal int64, ok bool) {
	a.mu.Lock()
    defer a.mu.Unlock()
    
    if !a.isOpen {
        return 0, false
    }


	return a.amount, true

    
}

func (a *Account) Deposit(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
    defer a.mu.Unlock()

	if !a.isOpen || a.amount + amt < 0{
        return 0, false
    }
    
	a.amount += amt

    return a.amount, true
    
}

func (a *Account) Withdraw(amt int64) (newBal int64, ok bool) {
	a.mu.Lock()
    defer a.mu.Unlock()

	if !a.isOpen || a.amount < amt || amt < 0 {
        return 0, false
    }


    a.amount -= amt

    return a.amount, true
    
}

func (a *Account) Close() (pay int64, ok bool) {
	a.mu.Lock()
    defer a.mu.Unlock()

	if !a.isOpen {
        return 0, false
    }


    a.isOpen = false
    return a.amount, true
}
