package service

import (
	"context"

	"tictactoe/internal/domain/models"
	"tictactoe/internal/ports"
)

type standardGameService struct {
	repo ports.StandardGameRepository
	ai   ports.AIService
}

// NewStandardGameService creates a new instance of the standard game service.
func NewStandardGameService(repo ports.StandardGameRepository, ai ports.AIService) ports.StandardGameService {
	return &standardGameService{
		repo: repo,
		ai:   ai,
	}
}

// CreateGame initializes a new 3x3 Tic-Tac-Toe session.
func (s *standardGameService) CreateGame(ctx context.Context) (string, error) {
	newGame := models.NewStandardGame()

	if err := s.repo.Save(ctx, newGame); err != nil {
		return "", err
	}

	return newGame.ID, nil
}

// GetGameState retrieves the current state of a standard game.
func (s *standardGameService) GetGameState(ctx context.Context, id string) (*models.StandardGame, error) {
	return s.repo.FindByID(ctx, id)
}

// MakeMove processes a player's move in a standard game.
func (s *standardGameService) MakeMove(ctx context.Context, id string, cellIdx int) (*models.StandardGame, error) {
	game, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	// Bouncers: Validate move preconditions
	if game.IsGameOver {
		return nil, models.ErrGameOver
	}

	if cellIdx < 0 || cellIdx > 8 {
		return nil, models.ErrInvalidMove
	}

	if game.Board[cellIdx] != models.Empty {
		return nil, models.ErrCellAlreadyTaken
	}

	// Execute Move
	game.Board[cellIdx] = game.CurrentPlayer

	// Check Winner
	if winner := models.CalculateWinner(game.Board[:]); winner != models.Empty {
		game.Winner = winner
		game.IsGameOver = true
	} else if models.IsBoardFull(game.Board[:]) {
		game.Winner = models.Tie
		game.IsGameOver = true
	}

	// Switch Player (only if game not over)
	if !game.IsGameOver {
		if game.CurrentPlayer == models.PlayerX {
			game.CurrentPlayer = models.PlayerO
		} else {
			game.CurrentPlayer = models.PlayerX
		}
	}

	if err := s.repo.Save(ctx, game); err != nil {
		return nil, err
	}

	return game, nil
}

// RequestBotMove asks the AI for a move and executes it.
func (s *standardGameService) RequestBotMove(ctx context.Context, id string) (*models.StandardGame, error) {
	game, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if game.IsGameOver {
		return nil, models.ErrGameOver
	}

	cellIdx, err := s.ai.GetStandardMove(ctx, game)
	if err != nil {
		return nil, err
	}

	return s.MakeMove(ctx, id, cellIdx)
}
