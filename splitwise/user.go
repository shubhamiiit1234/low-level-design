package main

import (
	"fmt"
)

type User struct {
	UserID   string
	Name     string
	Email    string
	Balances map[string]float64 // u owe to other user
}

func NewUser(id, name, email string) *User {
	return &User{
		UserID:   id,
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
	fmt.Println(u.UserID, " Paid ", amountToBePaid, " to ", receiverUserId)
}

func (u *User) ViewBalances() {
	for userID, balance := range u.Balances {
		fmt.Println(u.UserID, " owes to ", userID, " ", balance, " Rs.")
	}
}
