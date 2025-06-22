package main

type SplitStrategy interface {
	ComputeSplits(amount float64, participants []*User) []Split
	UpdateExpense(paidBy string, participants []*User, expense *Expense)
}
