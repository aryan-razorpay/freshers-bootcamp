package main

import (
	"fmt"
	"sync"
)

type BankAccount struct {
	balance int
	mu      sync.Mutex
}

func (account *BankAccount) Deposit(amount int) {
	account.mu.Lock()
	defer account.mu.Unlock()
	account.balance += amount
	fmt.Printf("Deposited Rs.%d, New balance: Rs.%d\n", amount, account.balance)
}

func (account *BankAccount) Withdraw(amount int) bool {
	account.mu.Lock()
	defer account.mu.Unlock()
	if account.balance < amount {
		fmt.Printf("Withdrawal of Rs.%d failed. Insufficient funds. Current balance: Rs.%d\n", amount, account.balance)
		return false
	}
	account.balance -= amount
	fmt.Printf("Withdrew Rs.%d, New balance: Rs.%d\n", amount, account.balance)
	return true
}

func main() {
	account := &BankAccount{balance: 500}

	var wg sync.WaitGroup

	wg.Add(4)

	go func() {
		defer wg.Done()
		account.Deposit(200)
	}()

	go func() {
		defer wg.Done()
		account.Withdraw(150)
	}()

	go func() {
		defer wg.Done()
		account.Deposit(300)
	}()

	go func() {
		defer wg.Done()
		account.Withdraw(700)
	}()

	wg.Wait()

	fmt.Printf("Final balance: Rs.%d\n", account.balance)
}
