package main

type Snake struct {
	SourcePosition int
	EndPosition    int
}

func NewSnake(sourcePosition, endPosition int) *Snake {
	return &Snake{
		SourcePosition: sourcePosition,
		EndPosition:    endPosition,
	}
}
