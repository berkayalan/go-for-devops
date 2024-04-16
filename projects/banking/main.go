package main

import (
	"fmt"
)

type Account struct {
	AccountOwner string
	TenantID     string
	AccountName  string
	AccountID    string
	IsActive     bool
	Currency     string
	Amount       float64
}

type Tenant struct {
	TenantID   string
	TenantName string
	IsActive   bool
}

type IAtm interface {
	WithdrawMoney() float64
	AddMoney() float64
	ShowDetails() string
}

func (a Account) WithdrawMoney() float64 {

	var amount float64

	fmt.Println("Please enter how much money you want to withdraw.")

	fmt.Scanln(&amount)

	rest_amount := a.Amount - amount

	fmt.Printf("%v %v left in your account.", rest_amount, a.Currency)

	return rest_amount

}

func (a Account) AddMoney() float64 {

	var amount float64

	fmt.Println("Please enter how much money you want to add.")

	fmt.Scanln(&amount)

	new_amount := a.Amount - amount

	fmt.Printf("New amount in your account is %v %v.", new_amount, a.Currency)

	return new_amount

}

func (a Account) ShowDetails() string {

	return "cxfsa"

}

func (a Tenant) ShowDetails() string {

	return "cxfsa"

}

func main() {

	/* a1 := Account{TenantID: "22975", AccountOwner: "Berkay Alan", AccountName: "Saving",
		AccountID: "1155", IsActive: true, Currency: "€", Amount: 5000}
	a2 := Account{TenantID: "22975", AccountOwner: "Berkay Alan", AccountName: "PocketMoney",
		AccountID: "1156", IsActive: true, Currency: "€", Amount: 1366.44}
	a3 := Account{TenantID: "66734", AccountOwner: "Evgeny Keck", AccountName: "Investment",
		AccountID: "6567", IsActive: true, Currency: "€", Amount: 17977}
	a4 := Account{TenantID: "66734", AccountOwner: "Evgeny Keck", AccountName: "Holiday",
		AccountID: "6578", IsActive: false, Currency: "€", Amount: 300}

	all_accounts := [4]Account{a1, a2, a3, a4}

	t1 := Tenant{TenantID: "22975", TenantName: "Berkay Alan", IsActive: true}
	t2 := Tenant{TenantID: "66734", TenantName: "Evgeny Keck", IsActive: true}
	t3 := Tenant{TenantID: "43254", TenantName: "Saang Kell", IsActive: false}

	all_tenants := [3]Tenant{t1, t2, t3}

	fmt.Printf("%v\n ", all_accounts)
	fmt.Printf("%v\n ", all_tenants) */

	transaction := GetTransaction()

	choose_count := 0

	for choose_count <= 3 {
		for transaction > 0 && transaction <= 5 {
			if transaction == 5 {
				choose_count += 1
				if choose_count < 3 {
					fmt.Println("You made a wrong choose! Please try again.")
					transaction = GetTransaction()
					continue
				} else if choose_count == 3 {
					fmt.Println("You made 3 wrong choose! Your account is suspected!")
					break
				}
				break
			} else {
				switch transaction {
				case 1:
					fmt.Println("1")
					continue
				case 2:
					fmt.Println("2")
				case 3:
					fmt.Println("3")
				case 4:
					break
				}
				break
			}
		}
	}
}

func GetTransaction() int {

	var transaction int

	fmt.Println("Please select what you want to do.\n1 : Show my user details \n2 : Withdraw Money\n3 : Add Money\n4 : Exit")

	fmt.Scanln(&transaction)
	if transaction > 0 && transaction <= 4 {
		return transaction
	} else {
		return 5 // Error Code
	}
}
