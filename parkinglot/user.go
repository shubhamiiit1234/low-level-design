package main

type User struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	Mobile   string `json:"mobile"`
}

type Admin struct {
	User `json:"user"`
}

type Customer struct {
	User    `json:"user"`
	Vehicle Vehicle `json:"vehicle"`
}
