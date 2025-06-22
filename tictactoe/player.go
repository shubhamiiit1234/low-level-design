package main

type Player struct {
	ID     int
	Symbol Symbol
}

func NewPlayer(id int, symbol Symbol) *Player {
	return &Player{
		ID:     id,
		Symbol: symbol,
	}
}
