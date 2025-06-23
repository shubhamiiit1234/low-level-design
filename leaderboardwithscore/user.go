package main

import "fmt"

type User struct {
	ID    string
	Score int
}

func NewUser(id string) *User {
	return &User{
		ID:    id,
		Score: 0,
	}
}

func (u *User) AddScore(score int) {
	u.Score += score
}

func (u *User) ResetScore() {
	u.Score = 0
}

func (u *User) ViewScore() {
	fmt.Println("user ", u.ID, "'s score is: ", u.Score)
}
