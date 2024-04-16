package main

import (
	"fmt"
)

type Account struct {
	AccountOwner string
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
	AccountID  array
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

	a1 := Account{AccountOwner: "Berkay Alan", AccountName: "Saving",
		AccountID: "1155", IsActive: true, Currency: "€", Amount: 5000}
	a2 := Account{AccountOwner: "Berkay Alan", AccountName: "PocketMoney",
		AccountID: "1156", IsActive: true, Currency: "€", Amount: 1366.44}
	a3 := Account{AccountOwner: "Evgeny Keck", AccountName: "Investment",
		AccountID: "6567", IsActive: true, Currency: "€", Amount: 17977}
	a4 := Account{AccountOwner: "Evgeny Keck", AccountName: "Holiday",
		AccountID: "6578", IsActive: false, Currency: "€", Amount: 300}

	all_accounts := [4]Account{a1, a2, a3, a4}

	t1 := Tenant{TenantID: "22975", TenantName: "Berkay Alan", IsActive: true, AccountID: [2]string{"1155", "1156"}}
	t2 := Tenant{TenantID: "66734", TenantName: "Evgeny Keck", IsActive: true, AccountID: [2]string{"6567", "6578"}}
	t3 := Tenant{TenantID: "43254", TenantName: "Saang Kell", IsActive: false, AccountID: [1]string{"3245"}}

	all_tenants := [3]Tenant{t1, t2, t3}

}
