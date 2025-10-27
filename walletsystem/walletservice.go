package main

import (
	"fmt"
	"sync"
)

type WalletService struct {
	Users map[string]*User // Stores all registered users
	mu    sync.Mutex       // Protects concurrent user creation
}

// NewWalletService initializes the wallet management service
func NewWalletService() *WalletService {
	return &WalletService{
		Users: make(map[string]*User),
	}
}

// CreateUser registers a new user and returns it
func (ws *WalletService) CreateUser(name string) *User {
	ws.mu.Lock()
	defer ws.mu.Unlock()

	user := NewUser(name)
	ws.Users[user.ID] = user
	fmt.Printf("User '%s' created with Wallet\n", name)
	return user
}

// TransferFunds safely transfers money between two user wallets
func (ws *WalletService) TransferFunds(fromID, toID string, amount float64) error {
	from := ws.Users[fromID]
	to := ws.Users[toID]

	if from == nil || to == nil {
		return fmt.Errorf("invalid user ID(s)")
	}

	return from.Wallet.Transfer(to.Wallet, amount)
}
