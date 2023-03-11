package main

import (
	"fmt"
	"go-poo/accounts"
	"go-poo/users"
)

const firstAccountBills = 200
const secondAccountBills = 400

func PayBill(account verifyAccount, billValue float64) float64 {
	account.Withdraw(billValue)
	return account.GetBalance()
}

type verifyAccount interface {
	Withdraw(value float64) string
	GetBalance() float64
}

func main() {
	fmt.Println()

	firstAccount := accounts.SavingAccount{
		Holder:    users.Holder{Name: "Joseph", Document: "000.111.222-33", Profession: "Software Developer"},
		Agency:    1010,
		Account:   95812,
		Operation: 2,
	}
	firstAccount.Deposit(1000)
	fmt.Println(firstAccount.Holder.Name, "has an account with $", firstAccount.GetBalance())
	fmt.Println(firstAccount.Holder.Name, "is paying bills on a value of", firstAccountBills)
	fmt.Println(firstAccount.Holder.Name, "account now has $", PayBill(&firstAccount, firstAccountBills))

	fmt.Println()

	secondAccount := accounts.CheckingAccount{
		Holder:  users.Holder{Name: "Charlotte", Document: "999.888.777-66", Profession: "Product Manager"},
		Agency:  1010,
		Account: 95812,
	}
	secondAccount.Deposit(1700)
	fmt.Println(secondAccount.Holder.Name, "has an account with $", secondAccount.GetBalance())
	fmt.Println(secondAccount.Holder.Name, "is paying bills on a value of $", secondAccountBills)
	fmt.Println(secondAccount.Holder.Name, "account now has $", PayBill(&secondAccount, secondAccountBills))

	fmt.Println()
}
