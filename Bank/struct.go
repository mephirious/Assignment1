package Bank

import "fmt"

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func NewBankAccount(accountNumber, holderName string) *BankAccount {
	return &BankAccount{AccountNumber: accountNumber, HolderName: holderName, Balance: 0}
}

func (b *BankAccount) Deposit(amount float64) {
	b.Balance += amount
}

func (b *BankAccount) Withdraw(amount float64) {
	b.Balance -= amount
}

func (b BankAccount) GetBalance() {
	fmt.Printf("Current balance: %.2f\n", b.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, v := range transactions {
		if v < 0 {
			account.Withdraw(v)
			continue
		}
		account.Deposit(v)
	}
}
