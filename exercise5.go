package main

import "fmt"

// Exercisse 5 create Bank account struct The function should return an updated account.

type Account struct {
	Owner   string
	Balance float64
}

func deposit(acc Account, amount float64) Account {
	acc.Balance += amount
	return acc
}
func main() {
	account := Account{Owner: "John Doe", Balance: 100.0}
	fmt.Printf("Initial account: %+v\n", account)

	account = deposit(account, 50.0)
	fmt.Printf("Updated account after deposit: %+v\n", account)
}
