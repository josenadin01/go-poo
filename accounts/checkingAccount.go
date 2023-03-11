package accounts

import "go-poo/users"

type CheckingAccount struct {
	Holder          users.Holder
	Agency, Account int
	balance         float64
}

// Ponteiro apontando para a conta que está chamando o método
// Será uma function que o struct terá acesso
func (account *CheckingAccount) Withdraw(withdrawalValue float64) string {
	canWithdraw := withdrawalValue > 0 && withdrawalValue <= account.balance

	if !canWithdraw {
		return "Insufficient funds"
	}

	account.balance -= withdrawalValue
	return "Successful withdrawal"
}

func (account *CheckingAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		account.balance += depositValue
		return "Deposito realizado com sucesso", account.balance
	} else {
		return "Valor do deposito menor que zero", account.balance
	}
}

func (account *CheckingAccount) Transfer(transferValue float64, destinyAccount *CheckingAccount) bool {
	if transferValue < account.balance && transferValue > 0 {
		account.balance -= transferValue
		destinyAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}

func (account *CheckingAccount) GetBalance() float64 {
	return account.balance
}
