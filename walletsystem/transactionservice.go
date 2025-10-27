package main

import "fmt"

type TransactionService struct{}

// GetTransactionHistory prints the list of all transactions for a wallet
func (ts *TransactionService) GetTransactionHistory(wallet *Wallet) {
	fmt.Printf("\nTransaction History for %s:\n", wallet.UserID)
	for _, t := range wallet.Transactions {
		fmt.Printf("➡️ [%s] ₹%.2f | %s | %s\n",
			t.Type, t.Amount, t.Note, t.Timestamp.Format("15:04:05"))
	}
}
