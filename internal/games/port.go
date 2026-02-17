package games

import "chess-console/internal/dto"

//go:generate mockery
type Games interface {
	SwitchTurn()
	IsGameOver() bool
	Move(sr, sc, er, ec int, turn dto.Color) error
	Print()
	IsValidMove(sr, sc, er, ec int, turn dto.Color) error
	GetTurn() dto.Color
}
