package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
)

type Wallet struct {
	UserID       string
	Balance      float64
	Transactions []Transaction
	mu           sync.Mutex
}

func NewWallet(userID string) *Wallet {
	return &Wallet{
		UserID:       userID,
		Balance:      0.0,
		Transactions: []Transaction{},
	}
}

// Credit adds money to the wallet
func (w *Wallet) Credit(amount float64, note string) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Balance += amount
	w.Transactions = append(w.Transactions, Transaction{
		ID:        uuid.New().String(),
		Type:      CREDIT,
		Amount:    amount,
		Timestamp: time.Now(),
		Note:      note,
	})
	fmt.Printf("[CREDIT] ₹%.2f credited to %s (New Balance: ₹%.2f)\n", amount, w.UserID, w.Balance)
}

// Debit deducts money from the wallet if sufficient balance is available
func (w *Wallet) Debit(amount float64, note string) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.Balance < amount {
		return fmt.Errorf("[ERROR] insufficient funds in wallet %s", w.UserID)
	}
	w.Balance -= amount
	w.Transactions = append(w.Transactions, Transaction{
		ID:        uuid.New().String(),
		Type:      DEBIT,
		Amount:    amount,
		Timestamp: time.Now(),
		Note:      note,
	})
	fmt.Printf("[DEBIT] ₹%.2f debited from %s (New Balance: ₹%.2f)\n", amount, w.UserID, w.Balance)
	return nil
}

// Transfer moves funds between wallets atomically
func (w *Wallet) Transfer(to *Wallet, amount float64) error {
	if err := w.Debit(amount, "Transfer to "+to.UserID); err != nil {
		return err
	}
	to.Credit(amount, "Transfer from "+w.UserID)
	fmt.Printf("[TRANSFER] ₹%.2f transferred from %s → %s\n", amount, w.UserID, to.UserID)
	return nil
}
