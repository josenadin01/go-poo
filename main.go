package main

import (
	"fmt"
	"go-poo/accounts"
)

func main() {
	contaDaSilvia := accounts.CheckingAccount{Holder: "Silvia", Balance: 300}
	contaDoGustavo := accounts.CheckingAccount{Holder: "Gustavo", Balance: 100}

	status := contaDoGustavo.Transfer(-200, &contaDaSilvia)

	fmt.Println(status)
	fmt.Println(contaDaSilvia)
	fmt.Println(contaDoGustavo)
}
