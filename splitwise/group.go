package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Group struct {
	GroupId      string
	Name         string
	Participants []*User
	Expenses     []*Expense
}

func NewGroup(name string, participants []*User) *Group {
	return &Group{
		GroupId:      uuid.NewString(),
		Name:         name,
		Participants: participants,
		Expenses:     []*Expense{},
	}
}

func (g *Group) AddMember(user User) {
	g.Participants = append(g.Participants, &user)
	fmt.Println("user ", user.Name, " added to the group")
}

func (g *Group) AddExpense(expense *Expense) {
	g.Expenses = append(g.Expenses, expense)
}
