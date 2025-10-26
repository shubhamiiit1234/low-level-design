package main

import (
	"fmt"

	"github.com/google/uuid"
)

type Expense struct {
	Id            string
	PaidBy        string
	Amount        float64
	Description   string
	Participants  []*User
	Splits        []Split
	SplitStrategy SplitStrategy
}

func NewExpense(paidBy, grouptId, description string, participants []*User, amount float64, splitStrategy SplitStrategy) *Expense {
	fmt.Println("split list: ", splitStrategy.ComputeSplits(amount, participants))
	return &Expense{
		Id:            uuid.NewString(),
		PaidBy:        paidBy,
		Amount:        amount,
		Description:   description,
		Participants:  participants,
		Splits:        splitStrategy.ComputeSplits(amount, participants),
		SplitStrategy: splitStrategy,
	}
}
