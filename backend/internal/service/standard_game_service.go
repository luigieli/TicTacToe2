package service

import (
	"context"
	"log"
	"math/rand"
	"time"

	"tictactoe/internal/domain/models"
	"tictactoe/internal/ports"
)

type standardGameService struct {
	repo ports.StandardGameRepository
	ai   ports.AIService
}

// NewStandardGameService creates a new instance of the standard game service.
func NewStandardGameService(repo ports.StandardGameRepository, ai ports.AIService) ports.StandardGameService {
	rand.Seed(time.Now().UnixNano())
	return &standardGameService{
		repo: repo,
		ai:   ai,
	}
}

// CreateGame initializes a new 3x3 Tic-Tac-Toe session.
func (s *standardGameService) CreateGame(ctx context.Context, mode string) (string, error) {
	m := models.StandardGameMode(mode)
	newGame := models.NewStandardGame(m)

	if err := s.repo.Save(ctx, newGame); err != nil {
		return "", err
	}

	log.Printf("SERVICE: Created game %s with mode %s", newGame.ID, newGame.Mode)
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

	if game.IsGameOver || game.Board[cellIdx] != models.Empty {
		return nil, models.ErrInvalidMove
	}

	// 1. Player Move
	s.applyMove(game, cellIdx)

	// 2. CPU Move if applicable
	isPVE := game.Mode == models.PVE_EASY || game.Mode == models.PVE_MEDIUM || game.Mode == models.PVE_HARD
	if isPVE && !game.IsGameOver && game.CurrentPlayer == models.PlayerO {
		s.applyCPUMove(ctx, game)
	}

	if err := s.repo.Save(ctx, game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s *standardGameService) applyMove(game *models.StandardGame, cellIdx int) {
	game.Board[cellIdx] = game.CurrentPlayer

	if winner := models.CalculateWinner(game.Board[:]); winner != models.Empty {
		game.Winner = winner
		game.IsGameOver = true
	} else if models.IsBoardFull(game.Board[:]) {
		game.Winner = models.Tie
		game.IsGameOver = true
	}

	if !game.IsGameOver {
		if game.CurrentPlayer == models.PlayerX {
			game.CurrentPlayer = models.PlayerO
		} else {
			game.CurrentPlayer = models.PlayerX
		}
	}
}

func (s *standardGameService) applyCPUMove(ctx context.Context, game *models.StandardGame) {
	var move int
	var err error

	switch game.Mode {
	case models.PVE_HARD:
		// Use AIService for Minimax
		move, err = s.ai.GetStandardMove(ctx, game)
		if err != nil {
			move = s.getRandomMove(game)
		}
	case models.PVE_MEDIUM:
		move = s.getMediumMove(game)
	default: // EASY
		move = s.getRandomMove(game)
	}

	if move != -1 {
		s.applyMove(game, move)
	}
}

func (s *standardGameService) getRandomMove(game *models.StandardGame) int {
	var emptyCells []int
	for i, cell := range game.Board {
		if cell == models.Empty {
			emptyCells = append(emptyCells, i)
		}
	}
	if len(emptyCells) == 0 {
		return -1
	}
	return emptyCells[rand.Intn(len(emptyCells))]
}

func (s *standardGameService) getMediumMove(game *models.StandardGame) int {
	// 1. Try to win
	if m := s.findWinningMove(game, models.PlayerO); m != -1 {
		return m
	}
	// 2. Block player
	if m := s.findWinningMove(game, models.PlayerX); m != -1 {
		return m
	}
	// 3. Fallback to random
	return s.getRandomMove(game)
}

func (s *standardGameService) findWinningMove(game *models.StandardGame, player models.CellState) int {
	for i := 0; i < 9; i++ {
		if game.Board[i] == models.Empty {
			game.Board[i] = player
			winner := models.CalculateWinner(game.Board[:])
			game.Board[i] = models.Empty
			if winner == player {
				return i
			}
		}
	}
	return -1
}
