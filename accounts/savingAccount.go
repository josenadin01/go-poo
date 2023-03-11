package accounts

import "go-poo/users"

type SavingAccount struct {
	Holder                     users.Holder
	Agency, Account, Operation int
	balance                    float64
}

func (account *SavingAccount) Withdraw(withdrawalValue float64) string {
	canWithdraw := withdrawalValue > 0 && withdrawalValue <= account.balance

	if !canWithdraw {
		return "Insufficient funds"
	}

	account.balance -= withdrawalValue
	return "Successful withdrawal"
}

func (account *SavingAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		account.balance += depositValue
		return "Deposito realizado com sucesso", account.balance
	} else {
		return "Valor do deposito menor que zero", account.balance
	}
}

func (account *SavingAccount) GetBalance() float64 {
	return account.balance
}
