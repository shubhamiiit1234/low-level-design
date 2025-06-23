package main

type Ladder struct {
	SourcePosition int
	EndPosition    int
}

func NewLadder(sourcePosition, endPosition int) *Ladder {
	return &Ladder{
		SourcePosition: sourcePosition,
		EndPosition:    endPosition,
	}
}
