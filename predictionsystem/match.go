package main

import "time"

type Match struct {
	ID       string
	TeamA    string
	TeamB    string
	Result   string
	StartAt  time.Time
	IsClosed bool
}
