package main

import "fmt"

type Board struct {
	size   int
	Matrix [][]string
}

func NewBoard(size int) *Board {

	board := Board{
		size: size,
		Matrix: [][]string{
			{".", ".", "."},
			{".", ".", "."},
			{".", ".", "."},
		},
	}

	return &board
}

// TODO
func (b *Board) ResetBoard() {}

func (b *Board) ShowBoard() {
	for _, val1 := range b.Matrix {
		for _, val2 := range val1 {
			fmt.Print(val2, " ")
		}
		fmt.Println()
	}
}
