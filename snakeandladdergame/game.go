package main

import (
	"fmt"
)

type Game struct {
	Board         *Board
	Players       map[int]*Player // playerID -> Player
	Die           *Die
	CurrentPlayer int
}

func NewGame() *Game {
	return &Game{
		Board:         &Board{},
		Players:       map[int]*Player{},
		Die:           &Die{},
		CurrentPlayer: 0,
	}
}

func (g *Game) AddBoard(board *Board) {
	g.Board = board
}

func (g *Game) AddPlayer(player *Player) {
	g.Players[player.ID] = player
}

func (g *Game) Start() {
	fmt.Println("Game started!!!")

	for {
		var input string
		fmt.Println("Player ", g.CurrentPlayer, "'s turn")
		fmt.Println("Press r and enter to play your move (r means roll the die)")
		fmt.Scanln(&input)

		if input != "r" {
			fmt.Println("Invalid operation!!!")
		} else {
			numberOnDie := g.Die.Roll()
			fmt.Println("Number: ", numberOnDie)
			currentPlayer := g.Players[g.CurrentPlayer]
			currentPosition := currentPlayer.Position
			afterMovingPosition := currentPosition + numberOnDie
			finalPosition := g.Board.GetNewPosition(afterMovingPosition)

			if finalPosition <= 100 {
				g.Players[g.CurrentPlayer].Position = finalPosition
			}

			fmt.Println("after rolling, Player ", g.CurrentPlayer, "'s current position: ", finalPosition)

			if finalPosition == 100 {
				fmt.Println("Player ", g.CurrentPlayer, " wins!!!")
				return
			}
			g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
			fmt.Println("-------------------------------------------------------------------------------------")
		}
	}
}
