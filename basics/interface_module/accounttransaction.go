package interface_module

import "fmt"

type Account struct {
	accountName     string
	accountOwner    string
	accountCurrency string
	accountBalance  float64
}

func (account *Account) Deposit(amount float64) {
	account.accountBalance += amount
	fmt.Printf("%.2f %s added to your %s account, new balance of your account is %.2f %s.\n",
		amount, account.accountCurrency, account.accountName, account.accountBalance, account.accountCurrency)
}

func (account *Account) Withdraw(amount float64) {
	account.accountBalance -= amount
	fmt.Printf("%.2f %s withdrawed from your %s account, new balance of your account is %.2f %s.\n",
		amount, account.accountCurrency, account.accountName, account.accountBalance, account.accountCurrency)
}

func AccountTransaction() {
	a := Account{"Saving", "Berkay Alan", "Euro", 300}
	fmt.Printf("Welcome to Gitbank Sir %s!\n", a.accountOwner)
	a.Deposit(45)
	a.Withdraw(74)
	fmt.Printf("The balance of your account is %.2f %s. See you next time!\n", a.accountBalance, a.accountCurrency)

}
