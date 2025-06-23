package main

import "math/rand"

type Die struct{}

func NewDie() *Die {
	return &Die{}
}

func (d *Die) Roll() int {
	return rand.Intn(6) + 1
}
