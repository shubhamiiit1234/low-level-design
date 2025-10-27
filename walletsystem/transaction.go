package main

import "time"

type TransactionType string

// Possible transaction types
const (
	CREDIT   TransactionType = "CREDIT"   // Money added to wallet
	DEBIT    TransactionType = "DEBIT"    // Money removed from wallet
	TRANSFER TransactionType = "TRANSFER" // Money transferred between wallets
)

type Transaction struct {
	ID        string
	Type      TransactionType
	Amount    float64
	Timestamp time.Time
	Note      string
}
