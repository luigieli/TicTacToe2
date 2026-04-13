package ports

import (
	"context"
	"tictactoe/internal/domain/models"
)

// AIService defines the interface for suggesting moves in different game modes.
type AIService interface {
	GetStandardMove(ctx context.Context, game *models.StandardGame) (int, error)
	GetUltimateMove(ctx context.Context, game *models.Game) (int, int, error) // boardIdx, cellIdx
}
