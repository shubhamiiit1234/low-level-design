package main

type Board struct {
	Ladders []*Ladder
	Snakes  []*Snake
	Cells   [101]int
}

func NewBoard() *Board {
	cells := [101]int{}
	for i := 0; i < len(cells); i++ {
		cells[i] = i
	}

	return &Board{
		Ladders: []*Ladder{},
		Snakes:  []*Snake{},
		Cells:   cells,
	}
}

func (b *Board) AddLadder(sourcePosition, endPosition int) {
	newLadder := NewLadder(sourcePosition, endPosition)
	b.Ladders = append(b.Ladders, newLadder)
}

func (b *Board) AddSnake(sourcePosition, endPosition int) {
	newSnake := NewSnake(sourcePosition, endPosition)
	b.Snakes = append(b.Snakes, newSnake)
}

func (b *Board) GetNewPosition(sourcePosition int) int {
	for _, ladder := range b.Ladders {
		if sourcePosition == ladder.SourcePosition {
			return ladder.EndPosition
		}
	}

	for _, snake := range b.Snakes {
		if sourcePosition == snake.SourcePosition {
			return snake.EndPosition
		}
	}
	return sourcePosition
}
