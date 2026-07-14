package main

import (
	"errors"
	"fmt"
)

type Account struct {
	Owner   string
	Balance float64
}

func (acc *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("Account must be greater than 0")
	}
	acc.Balance += amount
	return nil
}

func (acc *Account) Withdraw(amount float64) error {
	if amount > acc.Balance {
		return errors.New("Insufficient funds")
	}
	acc.Balance -= amount
	return nil
}
func main() {
	acc := Account{
		Owner:   "Harmony",
		Balance: 200,
	}

	fmt.Println("\n--- Deposits ---")
	err := acc.Deposit(500)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Deposited ksh 500.00, New Balance: ksh %.2f\n", acc.Balance)
	}

	fmt.Println("\n--- Withdraw ---")
	err = acc.Withdraw(500)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Withdraw ksh 500.00, New Balance: ksh %.2f\n", acc.Balance)
	}
}
