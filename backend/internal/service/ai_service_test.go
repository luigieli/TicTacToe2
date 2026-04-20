package service

import (
	"context"
	"testing"
	"tictactoe/internal/domain/models"
)

func TestStandardAI_SuggestMove_ValidCell(t *testing.T) {
	// Arrange
	ai := NewAIService() // Should fail to compile/run as it's not implemented yet
	ctx := context.Background()

	game := &models.StandardGame{
		Board: [9]models.CellState{
			models.PlayerX, models.Empty,   models.PlayerO,
			models.Empty,   models.PlayerX, models.Empty,
			models.Empty,   models.Empty,   models.PlayerO,
		},
		CurrentPlayer: models.PlayerO,
		IsGameOver:    false,
	}

	// Act
	move, err := ai.GetStandardMove(ctx, game)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if game.Board[move] != models.Empty {
		t.Errorf("expected move to be to an empty cell, but cell %d is %s", move, game.Board[move])
	}
}

func TestStandardAI_BlocksWinningMove(t *testing.T) {
	// Arrange
	ai := NewAIService()
	ctx := context.Background()

	game := &models.StandardGame{
		Board: [9]models.CellState{
			models.PlayerX, models.Empty,   models.Empty,
			models.PlayerX, models.PlayerO, models.Empty,
			models.Empty,   models.Empty,   models.Empty, // X can win at 6
		},
		CurrentPlayer: models.PlayerO,
		IsGameOver:    false,
	}

	// Act
	move, err := ai.GetStandardMove(ctx, game)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if move != 6 {
		t.Errorf("expected AI to block at 6, but got %d", move)
	}
}
