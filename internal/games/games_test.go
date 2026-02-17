package games

import (
	"chess-console/internal/dto"
	"io"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)
	assert.NotNil(t, game)
	assert.Equal(t, dto.White, game.GetTurn())
	assert.False(t, game.IsGameOver())
}

func TestSwitchTurn(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)
	assert.Equal(t, dto.White, game.GetTurn())
	game.SwitchTurn()
	assert.Equal(t, dto.Black, game.GetTurn())
	game.SwitchTurn()
	assert.Equal(t, dto.White, game.GetTurn())
}

func TestMove_Pawn(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)

	// Valid White Pawn move
	err := game.Move(6, 0, 5, 0, dto.White)
	assert.NoError(t, err)

	// Valid Black Pawn move
	game.SwitchTurn()
	err = game.Move(1, 0, 2, 0, dto.Black)
	assert.NoError(t, err)
}

func TestMove_InvalidTurn(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)

	// White turn, but trying to move a Black piece (at row 1)
	err := game.Move(1, 0, 2, 0, dto.White)
	assert.Error(t, err)
	assert.Equal(t, "not your turn", err.Error())
}

func TestMove_OutOfBounds(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)

	err := game.Move(6, 0, 8, 0, dto.White)
	assert.Error(t, err)
	assert.Equal(t, "out of bounds", err.Error())
}

func TestMove_NoPiece(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)

	err := game.Move(4, 4, 3, 4, dto.White)
	assert.Error(t, err)
	assert.Equal(t, "no piece at source", err.Error())
}

func TestMove_CaptureOwnPiece(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)

	// Pawn tries to move to a square occupied by a Rook of the same color
	err := game.Move(6, 0, 7, 0, dto.White)
	assert.Error(t, err)
	assert.Equal(t, "cannot capture own piece", err.Error())
}

func TestMove_RookBlocked(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)

	// Rook tries to move through its own pawn
	err := game.Move(7, 0, 5, 0, dto.White)
	assert.Error(t, err)
	assert.Equal(t, "path blocked", err.Error())
}

func TestMove_KnightValid(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)

	// Knight jumps over pawns
	err := game.Move(7, 1, 5, 2, dto.White)
	assert.NoError(t, err)
}

func TestIsGameOver(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger)
	service := game.(*Service)

	// Capture White King
	service.Grid[7][4] = nil
	assert.True(t, game.IsGameOver())
}

func TestService_ValidatePawn(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger).(*Service)

	// White pawn at (6,0)
	pawn := game.Grid[6][0]

	// Valid forward move
	assert.NoError(t, game.validatePawn(6, 0, 5, 0, pawn))

	// Invalid horizontal move
	assert.Error(t, game.validatePawn(6, 0, 6, 1, pawn))

	// Diagonal capture (no piece at target)
	assert.Error(t, game.validatePawn(6, 0, 5, 1, pawn))

	// Diagonal capture (piece at target)
	game.Grid[5][1] = &dto.Piece{Type: dto.Pawn, Color: dto.Black}
	assert.NoError(t, game.validatePawn(6, 0, 5, 1, pawn))
}

func TestService_ValidateRook(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger).(*Service)

	// Rook at (7,0)
	// Clear the path
	game.Grid[6][0] = nil

	// Valid vertical move
	assert.NoError(t, game.validateRook(7, 0, 5, 0))

	// Valid horizontal move
	game.Grid[7][0] = nil
	game.Grid[4][4] = &dto.Piece{Type: dto.Rook, Color: dto.White}
	assert.NoError(t, game.validateRook(4, 4, 4, 7))

	// Blocked path
	game.Grid[4][6] = &dto.Piece{Type: dto.Pawn, Color: dto.White}
	assert.Error(t, game.validateRook(4, 4, 4, 7))
}

func TestService_ValidateKnight(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger).(*Service)

	// Knight at (7,1)
	assert.NoError(t, game.validateKnight(7, 1, 5, 2))
	assert.NoError(t, game.validateKnight(7, 1, 5, 0))
	assert.Error(t, game.validateKnight(7, 1, 4, 1)) // Invalid vertical
}

func TestService_ValidateBishop(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger).(*Service)

	// Bishop at (7,2)
	game.Grid[6][3] = nil // Clear path
	assert.NoError(t, game.validateBishop(7, 2, 5, 4))

	// Blocked path
	game.Grid[6][3] = &dto.Piece{Type: dto.Pawn, Color: dto.White}
	assert.Error(t, game.validateBishop(7, 2, 5, 4))

	// Invalid horizontal move
	assert.Error(t, game.validateBishop(7, 2, 7, 5))
}

func TestService_ValidateQueen(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger).(*Service)

	// Queen at (7,3)
	game.Grid[7][3] = &dto.Piece{Type: dto.Queen, Color: dto.White}
	game.Grid[6][3] = nil // Clear vertical
	game.Grid[6][4] = nil // Clear diagonal
	game.Grid[7][4] = nil // Clear horizontal

	// Horizontal
	assert.NoError(t, game.validateQueen(7, 3, 7, 5))
	// Vertical
	assert.NoError(t, game.validateQueen(7, 3, 5, 3))
	// Diagonal
	assert.NoError(t, game.validateQueen(7, 3, 5, 5))

	// Invalid move
	assert.Error(t, game.validateQueen(7, 3, 5, 4))
}

func TestService_ValidateKing(t *testing.T) {
	logger := slog.New(slog.NewTextHandler(io.Discard, nil))
	game := NewGame(logger).(*Service)

	// King at (7,4)
	game.Grid[6][4] = nil // Clear forward

	// Valid 1 step moves
	assert.NoError(t, game.validateKing(7, 4, 6, 4))
	assert.NoError(t, game.validateKing(7, 4, 6, 5))
	assert.NoError(t, game.validateKing(7, 4, 7, 5))

	// Invalid 2 step move
	assert.Error(t, game.validateKing(7, 4, 5, 4))
}
