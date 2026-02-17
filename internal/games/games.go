package games

import (
	"chess-console/internal/dto"
	"chess-console/pkg/shared/utils"
	"errors"
	"fmt"
	"log/slog"
)

type Service struct {
	Turn   dto.Color
	Grid   [8][8]*dto.Piece
	Logger *slog.Logger
}

func NewGame(logger *slog.Logger) Games {
	s := &Service{
		Logger: logger,
		Turn:   dto.White,
	}
	s.initialize()
	return s
}

func (s *Service) initialize() {
	// Pawns
	for i := 0; i < 8; i++ {
		s.Grid[1][i] = &dto.Piece{Type: dto.Pawn, Color: dto.Black}
		s.Grid[6][i] = &dto.Piece{Type: dto.Pawn, Color: dto.White}
	}

	setup := []dto.PieceType{dto.Rook, dto.Knight, dto.Bishop, dto.Queen, dto.King, dto.Bishop, dto.Knight, dto.Rook}

	for i, p := range setup {
		s.Grid[0][i] = &dto.Piece{Type: p, Color: dto.Black}
		s.Grid[7][i] = &dto.Piece{Type: p, Color: dto.White}
	}
}

func (g *Service) SwitchTurn() {
	if g.Turn == dto.White {
		g.Turn = dto.Black
	} else {
		g.Turn = dto.White
	}
}

func (g *Service) IsGameOver() bool {
	whiteKing, blackKing := false, false

	for i := 0; i < 8; i++ {
		for j := 0; j < 8; j++ {
			p := g.Grid[i][j]
			if p != nil && p.Type == dto.King {
				if p.Color == dto.White {
					whiteKing = true
				}
				if p.Color == dto.Black {
					blackKing = true
				}
			}
		}
	}
	return !(whiteKing && blackKing)
}

func (s *Service) Print() {
	fmt.Println("  a s c d e f g h")
	for i := 0; i < 8; i++ {
		fmt.Print(8-i, " ")
		for j := 0; j < 8; j++ {
			if s.Grid[i][j] == nil {
				fmt.Print(". ")
			} else {
				fmt.Print(string(s.Grid[i][j].Type), " ")
			}
		}
		fmt.Println()
	}
}

func (s *Service) IsValidMove(sr, sc, er, ec int, turn dto.Color) error {

	if !inBounds(sr, sc) || !inBounds(er, ec) {
		return errors.New("out of bounds")
	}

	piece := s.Grid[sr][sc]
	if piece == nil {
		return errors.New("no piece at source")
	}

	if piece.Color != turn {
		return errors.New("not your turn")
	}

	dest := s.Grid[er][ec]
	if dest != nil && dest.Color == piece.Color {
		return errors.New("cannot capture own piece")
	}

	switch piece.Type {
	case dto.Pawn:
		return s.validatePawn(sr, sc, er, ec, piece)
	case dto.Rook:
		return s.validateRook(sr, sc, er, ec)
	case dto.Knight:
		return s.validateKnight(sr, sc, er, ec)
	case dto.Bishop:
		return s.validateBishop(sr, sc, er, ec)
	case dto.Queen:
		return s.validateQueen(sr, sc, er, ec)
	case dto.King:
		return s.validateKing(sr, sc, er, ec)
	}

	return nil
}

func inBounds(r, c int) bool {
	return r >= 0 && r < 8 && c >= 0 && c < 8
}

func (s *Service) validatePawn(sr, sc, er, ec int, p *dto.Piece) error {
	dir := -1
	if p.Color == dto.Black {
		dir = 1
	}

	if sc == ec && s.Grid[er][ec] == nil && er == sr+dir {
		return nil
	}

	if utils.Abs(sc-ec) == 1 && er == sr+dir && s.Grid[er][ec] != nil {
		return nil
	}

	return errors.New("invalid pawn move")
}

func (s *Service) validateRook(sr, sc, er, ec int) error {
	if sr != er && sc != ec {
		return errors.New("invalid rook move")
	}
	return s.isPathClear(sr, sc, er, ec)
}

func (s *Service) isPathClear(sr, sc, er, ec int) error {
	dr := utils.Sign(er - sr)
	dc := utils.Sign(ec - sc)

	r, c := sr+dr, sc+dc

	for r != er || c != ec {
		if s.Grid[r][c] != nil {
			return errors.New("path blocked")
		}
		r += dr
		c += dc
	}
	return nil
}

func (s *Service) Move(sr, sc, er, ec int, turn dto.Color) error {
	if err := s.IsValidMove(sr, sc, er, ec, turn); err != nil {
		return err
	}

	s.Grid[er][ec] = s.Grid[sr][sc]
	s.Grid[sr][sc] = nil
	return nil
}

func (s *Service) validateKnight(sr, sc, er, ec int) error {
	dr := utils.Abs(er - sr)
	dc := utils.Abs(ec - sc)

	if (dr == 2 && dc == 1) || (dr == 1 && dc == 2) {
		return nil
	}

	return errors.New("invalid knight move")
}

func (s *Service) validateBishop(sr, sc, er, ec int) error {
	dr := utils.Abs(er - sr)
	dc := utils.Abs(ec - sc)

	if dr != dc {
		return errors.New("invalid bishop move")
	}

	return s.isPathClear(sr, sc, er, ec)
}

func (s *Service) validateQueen(sr, sc, er, ec int) error {

	// Diagonal → Bishop logic
	if utils.Abs(er-sr) == utils.Abs(ec-sc) {
		return s.validateBishop(sr, sc, er, ec)
	}

	// Horizontal / Vertical → Rook logic
	if sr == er || sc == ec {
		return s.validateRook(sr, sc, er, ec)
	}

	return errors.New("invalid queen move")
}

func (s *Service) validateKing(sr, sc, er, ec int) error {
	dr := utils.Abs(er - sr)
	dc := utils.Abs(ec - sc)

	if dr <= 1 && dc <= 1 {
		return nil
	}

	return errors.New("invalid king move")
}

func (s *Service) GetTurn() dto.Color {
	return s.Turn
}
