package main

import (
	"fmt"
)

type TicTacToe struct {
	Players    []*Player
	Board      *Board
	Turn       int
	TotalMoves int
}

func NewTicTacToe() *TicTacToe {
	return &TicTacToe{
		Players:    []*Player{},
		Board:      &Board{},
		Turn:       1,
		TotalMoves: 0,
	}
}

func (t *TicTacToe) AddPlayer(player []*Player) {
	t.Players = player
}

func (t *TicTacToe) InitializeBoard() {
	board := NewBoard(3)
	t.Board = board
}

func (t *TicTacToe) Start(turn int) {

	t.ViewBoard()

	for {
		var row, col int
		fmt.Println("Player ", turn, "'s turn")
		fmt.Print("Enter row and col with spaces: ")
		fmt.Scanln(&row, &col)
		// fmt.Println("row: ", row, " col: ", col)

		if t.Validate(row, col) {
			symbol := GetSymbol(turn)
			t.PlayAMove(row, col, symbol)
			t.ViewBoard()
			t.TotalMoves++

			state := t.CheckState(row, col, turn, t.TotalMoves)
			fmt.Println("state: ", state)
			switch state {
			case Player1Wins:
				fmt.Println("Player 1 wins!!!")
				return
			case Player0Wins:
				fmt.Println("Player 0 wins!!!")
				return
			case Draw:
				fmt.Println("Match Draw!!!")
				return
			}
			turn = (turn + 1) % 2
		} else {
			fmt.Println("Invalid position")
		}
		// t.ViewBoard()
	}
}

func GetSymbol(turn int) Symbol {
	if turn == 1 {
		return X
	}

	return O
}

func (t *TicTacToe) ViewBoard() {
	t.Board.ShowBoard()
}

func (t *TicTacToe) PlayAMove(row, col int, symbol Symbol) {
	t.Board.Matrix[row][col] = string(symbol)
}

func (t *TicTacToe) Validate(row, col int) bool {
	size := t.Board.size
	if row < 0 || row >= size || col < 0 || col >= size || t.Board.Matrix[row][col] != "." {
		return false
	}
	return true
}

func (t *TicTacToe) CheckState(row, col, turn, totalMoves int) string {
	symbol := t.Board.Matrix[row][col]
	size := t.Board.size

	win := true

	// Column check
	for i := 0; i < size; i++ {
		if t.Board.Matrix[i][col] != symbol {
			win = false
			break
		}
	}

	if win {
		switch turn {
		case 1:
			fmt.Println("line 112")
			return Player1Wins
		case 0:
			return Player0Wins
		}
	}

	win = true

	// Row check
	for i := 0; i < size; i++ {
		if t.Board.Matrix[row][i] != symbol {
			win = false
			break
		}
	}

	if win {
		switch turn {
		case 1:
			fmt.Println("line 132")
			return Player1Wins
		case 0:
			return Player0Wins
		}
	}

	win = true

	// Main Diagonal
	if row == col {
		for i := 0; i < size; i++ {
			if t.Board.Matrix[i][i] != symbol {
				win = false
				break
			}
		}

		if win {
			switch turn {
			case 1:
				fmt.Println("line 154")
				return Player1Wins
			case 0:
				return Player0Wins
			}
		}
	}

	win = true

	// Anti-Diagonal
	if row+col == size-1 {
		for i := 0; i < size; i++ {
			if t.Board.Matrix[i][size-1-i] != symbol {
				win = false
				break
			}
		}

		if win {
			switch turn {
			case 1:
				fmt.Println("line 176")
				return Player1Wins
			case 0:
				return Player0Wins
			}
		}
	}

	if totalMoves == size*size {
		return Draw
	}

	return Ongoing
}

/*
	. X .
	. X O
	. X O
*/
