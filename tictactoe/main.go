package main

import "fmt"

/*
	Requirements:
		1. 3x3 board
		2. Player takes turn by marking 'X' or 'O' on board
		3. Player wins if their symbols are in a row (Horizontally | Vertically | Diagonally)
		4. If board is filled, and no one wins => Draw
		5. Validation of moves
		6. Detect and announce the winner | Draw

	Core Entities:
		1. TicTacToe (Central class)
		2. Board (3x3)
		3. Symbol (enum) - 'X' | 'O'
		4. Player

	Fields/Methods:
		1. TicTacToe
			- []Player, [][]Board, Turn/CurrentPlayer
			- Start(turn), AddPlayer(player []*Player), InitializeBoard(), CheckState(board, row, col)
		2. Board
			- Size
			- NewBoard(), ResetBoard(board [][]string)
		3. Player
			- ID, Symbol
			- NewPlayer()

	Happy Flow:
		Initialize Board -> Add Players -> Start Game -> Player made a move -> Check State -> if Won -> Announce and Reset Board
																^					|
																|					|-> else if Draw -> Announce Draw and Reset Board
																|					|
																--------------------|
																		else

*/

func main() {
	fmt.Println("tic-tac-toe lld!!!")
	ticTacToe := NewTicTacToe()
	ticTacToe.InitializeBoard()

	player1 := &Player{
		ID:     1,
		Symbol: X,
	}

	player0 := &Player{
		ID:     0,
		Symbol: O,
	}

	ticTacToe.AddPlayer([]*Player{player1, player0})
	ticTacToe.Start(1)
}
