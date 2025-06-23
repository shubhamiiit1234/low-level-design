package main

import "fmt"

/*
	Requirements:
		1. The game should be played on a board with numbered cells (1 -> 100)
		2. Predefined set of Snakes and Ladders, connecting two cells
		3. Should support multiple players
		4. Move by rolling a die (random number between [1,6])
		5. Climb up/Slide down to the top of the ladder/tail of the snake if player comes on the source cell of any of these two
		6. Game should continue until one of the player reaches 100
		7. The game should handle multiple game sessions concurrently, allowing different groups of players to play independently (TBD)

	Core Entities:
		1. Dice
			- Roll() -> gives random number b/w [1,6]
		2. Player
			- ID, Name, position (int), symbol
			- Move()
		3. Board
			- []*Ladder, []*Snake, [101]int
			- GetNewPosition(sourcePosition)
		4. Ladder & Snake
			- SourcePosition, EndPosition
		5. Game
			- Board, map[string]*Player (playerID->Player), Die, currentPlayer
			- Start()

*/

func main() {
	fmt.Println("Snake and Ladder Game LLD!!!")

	game := NewGame()

	board := NewBoard()

	board.AddLadder(4, 100)
	board.AddSnake(13, 5)

	game.AddBoard(board)

	player1 := NewPlayer(0, "Tyler", "yellow")
	player2 := NewPlayer(1, "Gamma", "green")

	game.AddPlayer(player1)
	game.AddPlayer(player2)

	game.Start()
}
