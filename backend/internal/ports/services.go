package ports

import (
	"context"
	"tictactoe/internal/domain/dto"
	"tictactoe/internal/domain/models"
)

// GameService defines the business operations for Tic-Tac-Toe-2.
type GameService interface {
	CreateGame(ctx context.Context) (string, error)
	GetGameState(ctx context.Context, id string) (*models.Game, error)
	MakeMove(ctx context.Context, req dto.MoveRequest) (*models.Game, error)
}

// StandardGameService defines the business operations for a standard 3x3 game.
type StandardGameService interface {
	CreateGame(ctx context.Context, mode string) (string, error)
	GetGameState(ctx context.Context, id string) (*models.StandardGame, error)
	MakeMove(ctx context.Context, id string, cellIdx int) (*models.StandardGame, error)
}
