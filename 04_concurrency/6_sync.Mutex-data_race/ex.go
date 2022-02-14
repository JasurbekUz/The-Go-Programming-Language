package main

import (
	"fmt"
	"sync"
	"time"
)

var m = &sync.Mutex{}

type Account struct {
	name    string
	balance float64
}

func (a *Account) Withdraw(amount float64) {
	fmt.Println(time.Now().UnixNano(), "With")
	m.Lock()
	defer m.Unlock()
	a.balance = a.balance - amount
}

func (a *Account) Deposit(amount float64) {
	fmt.Println(time.Now().UnixNano(), "Dep")
	m.Lock()
	a.balance = a.balance + amount
	m.Unlock()
}

func (a *Account) GetBalance() float64 {
	m.Lock()
	defer m.Unlock()
	return a.balance
}

func main() {

	account := Account{name: "Muhammad"}

	go account.Withdraw(5000)

	go account.Deposit(10_000)

	fmt.Println(account.GetBalance())
}
