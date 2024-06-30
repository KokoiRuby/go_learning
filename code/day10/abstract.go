package main

import "fmt"

type Account struct {
	AccountNo string
	Password  string
	Balance   float64
}

func (acc *Account) Deposit(amount float64, pwd string) {
	if pwd != acc.Password {
		panic("Wrong password")
	}
	if amount <= 0 {
		fmt.Println("Invalid amount")
	}
	acc.Balance += amount
	fmt.Println("Deposit successful.")
}

func (acc *Account) Withdraw(amount float64, pwd string) {
	if pwd != acc.Password {
		panic("Wrong password")
	}
	if amount <= 0 || amount > acc.Balance {
		fmt.Println("Invalid amount")
	}
	acc.Balance -= amount
	fmt.Println("Withdraw successful.")
}

func (acc *Account) Query(pwd string) {
	if pwd != acc.Password {
		panic("Wrong password")
	}
	fmt.Printf("Balance is: %v\n", acc.Balance)

}

func main() {
	account := Account{
		AccountNo: "123456",
		Password:  "123456",
		Balance:   10000,
	}
	account.Query("123456")
	account.Withdraw(500, "123456")
	account.Query("123456")
	account.Deposit(1000, "123456")
	account.Query("123456")
}
