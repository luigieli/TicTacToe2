package service

import (
	"context"
	"errors"
	"tictactoe/internal/domain/models"
	"tictactoe/internal/ports"
)

type aiService struct{}

// NewAIService returns a new instance of AIService.
func NewAIService() ports.AIService {
	return &aiService{}
}

// GetStandardMove suggests a move for a 3x3 game.
func (s *aiService) GetStandardMove(_ context.Context, game *models.StandardGame) (int, error) {
	if game.IsGameOver {
		return -1, models.ErrGameOver
	}

	// Simple strategy: first available cell
	for i, cell := range game.Board {
		if cell == models.Empty {
			return i, nil
		}
	}

	return -1, errors.New("no empty cells available")
}

// GetUltimateMove suggests a move for a 9x9 game.
func (s *aiService) GetUltimateMove(_ context.Context, game *models.Game) (int, int, error) {
	if game.IsGameOver {
		return -1, -1, models.ErrGameOver
	}

	// To be implemented in later cycles
	return -1, -1, errors.New("not implemented")
}
