package main

import "github.com/google/uuid"

type User struct {
	ID     string
	Name   string
	Wallet *Wallet
}

func NewUser(name string) *User {
	return &User{
		ID:     uuid.NewString(),
		Name:   name,
		Wallet: NewWallet(name),
	}
}
