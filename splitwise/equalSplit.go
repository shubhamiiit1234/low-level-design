package main

type EqualSplit struct{}

func (e *EqualSplit) ComputeSplits(amount float64, participants []*User) []Split {
	totalParticipants := len(participants)
	perPersonSplitAmount := amount / float64(totalParticipants)
	splits := []Split{}
	for _, participant := range participants {
		splits = append(splits, Split{
			UserID: participant.ID,
			Amount: perPersonSplitAmount,
		})
	}
	return splits
}

func (e *EqualSplit) UpdateExpense(paidBy string, participants []*User, expense *Expense) {
	for _, participant := range participants {
		if participant.ID != paidBy {
			participant.Balances[paidBy] += expense.Splits[0].Amount
		}
	}
}
