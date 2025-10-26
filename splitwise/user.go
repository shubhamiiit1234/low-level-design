package main

import (
	"fmt"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Balances map[string]float64 // u owe to other user
}

func NewUser(id, name, email string) *User {
	return &User{
		ID:       id,
		Name:     name,
		Email:    email,
		Balances: make(map[string]float64),
	}
}

func (u *User) SettleBalance(expense *Expense) {
	amountToBePaid := u.Balances[expense.PaidBy]
	receiverUserId := expense.PaidBy
	// Logic to pay
	u.Balances[receiverUserId] -= amountToBePaid
	fmt.Println(u.ID, " Paid ", amountToBePaid, " to ", receiverUserId)
}

func (u *User) ViewBalances() {
	for userID, balance := range u.Balances {
		fmt.Println(u.ID, " owes to ", userID, " ", balance, " Rs.")
	}
}
