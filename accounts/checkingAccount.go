package accounts

type CheckingAccount struct {
	Holder  string
	Agency  int
	Account int
	Balance float64
}

// Ponteiro para conta que está chamando
// Será uma function que o struct terá acesso
func (account *CheckingAccount) Withdraw(withdrawalValue float64) string {
	canWithdraw := withdrawalValue > 0 && withdrawalValue <= account.Balance

	if !canWithdraw {
		return "Insufficient funds"
	}

	account.Balance -= withdrawalValue
	return "Successful withdrawal"
}

func (account *CheckingAccount) Deposit(depositValue float64) (string, float64) {
	if depositValue > 0 {
		account.Balance += depositValue
		return "Deposito realizado com sucesso", account.Balance
	} else {
		return "Valor do deposito menor que zero", account.Balance
	}
}

func (account *CheckingAccount) Transfer(transferValue float64, destinyAccount *CheckingAccount) bool {
	if transferValue < account.Balance && transferValue > 0 {
		account.Balance -= transferValue
		destinyAccount.Deposit(transferValue)
		return true
	} else {
		return false
	}
}
