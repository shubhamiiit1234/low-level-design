// Final Go implementation of Chess with Move tracking, Check & Checkmate
package chess

import (
	"fmt"
)

// --- Enums & Helpers ---
type Color string

const (
	White Color = "White"
	Black Color = "Black"
)

type Position struct {
	X, Y int
}

func (p Position) Equals(other Position) bool {
	return p.X == other.X && p.Y == other.Y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// --- Piece Interface & Implementations ---
type Piece interface {
	Color() Color
	Name() string
	IsValidMove(board *Board, from, to Position) bool
}

type BasePiece struct {
	color Color
}

func (bp *BasePiece) Color() Color {
	return bp.color
}

// Rook
type Rook struct{ BasePiece }

func NewRook(color Color) *Rook {
	return &Rook{BasePiece{color}}
}

func (r *Rook) Name() string { return "Rook" }

func (r *Rook) IsValidMove(board *Board, from, to Position) bool {
	return from.X == to.X || from.Y == to.Y
}

// King
type King struct{ BasePiece }

func NewKing(color Color) *King {
	return &King{BasePiece{color}}
}

func (k *King) Name() string { return "King" }

func (k *King) IsValidMove(board *Board, from, to Position) bool {
	dx := abs(from.X - to.X)
	dy := abs(from.Y - to.Y)
	return dx <= 1 && dy <= 1
}

// --- Board & Cell ---
type Cell struct {
	Position Position
	Piece    Piece
}

type Board struct {
	Cells [8][8]*Cell
}

func NewBoard() *Board {
	b := &Board{}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			b.Cells[i][j] = &Cell{Position: Position{i, j}}
		}
	}
	b.initializePieces()
	return b
}

func (b *Board) GetCell(pos Position) *Cell {
	if pos.X < 0 || pos.X >= 8 || pos.Y < 0 || pos.Y >= 8 {
		return nil
	}
	return b.Cells[pos.X][pos.Y]
}

func (b *Board) FindKing(color Color) *Cell {
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			cell := b.Cells[i][j]
			if cell.Piece != nil && cell.Piece.Name() == "King" && cell.Piece.Color() == color {
				return cell
			}
		}
	}
	return nil
}

func (b *Board) initializePieces() {
	b.Cells[0][0].Piece = NewRook(White)
	b.Cells[0][7].Piece = NewRook(White)
	b.Cells[7][0].Piece = NewRook(Black)
	b.Cells[7][7].Piece = NewRook(Black)
	b.Cells[0][4].Piece = NewKing(White)
	b.Cells[7][4].Piece = NewKing(Black)
}

// --- Player ---
type Player struct {
	Color Color
}

// --- Move ---
type Move struct {
	From          Position
	To            Position
	MovedPiece    Piece
	CapturedPiece Piece
}

// --- Game ---
type Game struct {
	Board         *Board
	WhitePlayer   *Player
	BlackPlayer   *Player
	CurrentPlayer *Player
	GameOver      bool
	MoveHistory   []*Move
}

func NewGame() *Game {
	board := NewBoard()
	white := &Player{Color: White}
	black := &Player{Color: Black}
	return &Game{
		Board:         board,
		WhitePlayer:   white,
		BlackPlayer:   black,
		CurrentPlayer: white,
	}
}

func (g *Game) MakeMove(from, to Position) (*Move, bool) {
	if g.GameOver {
		fmt.Println("Game is already over.")
		return nil, false
	}

	fromCell := g.Board.GetCell(from)
	toCell := g.Board.GetCell(to)

	if fromCell == nil || toCell == nil || fromCell.Piece == nil {
		return nil, false
	}

	movingPiece := fromCell.Piece
	if movingPiece.Color() != g.CurrentPlayer.Color {
		return nil, false
	}

	if !movingPiece.IsValidMove(g.Board, from, to) {
		return nil, false
	}

	capturedPiece := toCell.Piece

	toCell.Piece = movingPiece
	fromCell.Piece = nil

	if g.IsKingInCheck(g.CurrentPlayer.Color) {
		fromCell.Piece = toCell.Piece
		toCell.Piece = capturedPiece
		return nil, false
	}

	move := &Move{
		From:          from,
		To:            to,
		MovedPiece:    movingPiece,
		CapturedPiece: capturedPiece,
	}
	g.MoveHistory = append(g.MoveHistory, move)

	// Switch player
	if g.CurrentPlayer == g.WhitePlayer {
		g.CurrentPlayer = g.BlackPlayer
	} else {
		g.CurrentPlayer = g.WhitePlayer
	}

	// Checkmate detection
	if g.IsCheckmate(g.CurrentPlayer.Color) {
		fmt.Printf("Checkmate! %s wins!\n", move.MovedPiece.Color())
		g.GameOver = true
	}

	return move, true
}

func (g *Game) IsKingInCheck(playerColor Color) bool {
	kingCell := g.Board.FindKing(playerColor)
	if kingCell == nil {
		return false
	}

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			cell := g.Board.Cells[i][j]
			if cell.Piece != nil && cell.Piece.Color() != playerColor {
				if cell.Piece.IsValidMove(g.Board, cell.Position, kingCell.Position) {
					return true
				}
			}
		}
	}
	return false
}

func (g *Game) IsCheckmate(playerColor Color) bool {
	if !g.IsKingInCheck(playerColor) {
		return false
	}
	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			from := g.Board.Cells[i][j]
			if from.Piece == nil || from.Piece.Color() != playerColor {
				continue
			}
			for x := 0; x < 8; x++ {
				for y := 0; y < 8; y++ {
					to := g.Board.Cells[x][y]
					if from.Position.Equals(to.Position) {
						continue
					}
					if !from.Piece.IsValidMove(g.Board, from.Position, to.Position) {
						continue
					}
					// Try move
					originalFrom := from.Piece
					originalTo := to.Piece

					to.Piece = from.Piece
					from.Piece = nil
					inCheck := g.IsKingInCheck(playerColor)
					from.Piece = originalFrom
					to.Piece = originalTo
					if !inCheck {
						return false
					}
				}
			}
		}
	}
	return true
}
