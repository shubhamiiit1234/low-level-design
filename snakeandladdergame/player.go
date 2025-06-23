package main

type Player struct {
	ID       int
	Name     string
	Position int
	Symbol   string
}

func NewPlayer(id int, name, symbol string) *Player {
	return &Player{
		ID:       id,
		Name:     name,
		Position: 0,
		Symbol:   symbol,
	}
}

func (p *Player) Move(position int) {
	p.Position = position
}
