package main

import (
	"fmt"
)

func main() {
	fmt.Println("wallet system LLD!!!")

	ws := NewWalletService()
	ts := &TransactionService{}

	alice := ws.CreateUser("Alice")
	bob := ws.CreateUser("Bob")

	alice.Wallet.Credit(1000, "Initial Deposit")
	bob.Wallet.Credit(500, "Initial Deposit")

	alice.Wallet.Debit(200, "Buy Item A")

	// Transfer money from Alice to Bob
	ws.TransferFunds(alice.ID, bob.ID, 300)

	fmt.Printf("\nFinal Balances:\nAlice: ₹%.2f\nBob: ₹%.2f\n",
		alice.Wallet.Balance, bob.Wallet.Balance)

	ts.GetTransactionHistory(alice.Wallet)
	ts.GetTransactionHistory(bob.Wallet)
}
