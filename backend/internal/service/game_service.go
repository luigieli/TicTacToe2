package service

import (
	"context"

	"github.com/google/uuid"
	"tictactoe/internal/domain/dto"
	"tictactoe/internal/domain/models"
	"tictactoe/internal/ports"
)

type gameService struct {
	repo ports.GameRepository
	ai   ports.AIService
}

// NewGameService creates a new instance of the game service.
func NewGameService(repo ports.GameRepository, ai ports.AIService) ports.GameService {
	return &gameService{
		repo: repo,
		ai:   ai,
	}
}

// CreateGame initializes a new Ultimate Tic Tac Toe session.
func (s *gameService) CreateGame(ctx context.Context, req dto.CreateGameRequest) (string, error) {
	mode := req.Mode
	if mode == "" {
		mode = models.ModePVP
	}

	newGame := &models.Game{
		ID:            uuid.New().String(),
		Mode:          mode,
		CurrentPlayer: models.PlayerX,
		NextBoardIdx:  -1,
		IsGameOver:    false,
	}

	// Initialize empty boards
	for i := 0; i < 9; i++ {
		newGame.SubBoards[i] = models.Board{
			Cells:  [9]models.CellState{},
			Winner: models.Empty,
		}
	}

	if err := s.repo.Save(ctx, newGame); err != nil {
		return "", err
	}

	return newGame.ID, nil
}

func (s *gameService) GetGameState(ctx context.Context, id string) (*models.Game, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *gameService) MakeMove(ctx context.Context, req dto.MoveRequest) (*models.Game, error) {
	game, err := s.repo.FindByID(ctx, req.GameID)
	if err != nil {
		return nil, err
	}

	// Bouncers: Validate move preconditions
	if game.IsGameOver {
		return nil, models.ErrGameOver
	}

	if game.NextBoardIdx != -1 && game.NextBoardIdx != req.BoardIdx {
		return nil, models.ErrWrongBoard
	}

	if req.BoardIdx < 0 || req.BoardIdx > 8 || req.CellIdx < 0 || req.CellIdx > 8 {
		return nil, models.ErrInvalidMove
	}

	targetBoard := &game.SubBoards[req.BoardIdx]
	if targetBoard.Winner != models.Empty {
		return nil, models.ErrBoardAlreadyWon
	}

	if targetBoard.Cells[req.CellIdx] != models.Empty {
		return nil, models.ErrCellAlreadyTaken
	}

	// Execute Move
	s.applyMove(game, req.BoardIdx, req.CellIdx)

	// If PVE mode, bot makes its move immediately
	if game.Mode == models.ModePVE && !game.IsGameOver {
		bIdx, cIdx, err := s.ai.GetUltimateMove(ctx, game)
		if err == nil {
			s.applyMove(game, bIdx, cIdx)
		}
	}

	if err := s.repo.Save(ctx, game); err != nil {
		return nil, err
	}

	return game, nil
}

func (s *gameService) applyMove(game *models.Game, boardIdx, cellIdx int) {
	targetBoard := &game.SubBoards[boardIdx]
	targetBoard.Cells[cellIdx] = game.CurrentPlayer

	// Check Sub-Board Winner
	if winner := models.CalculateWinner(targetBoard.Cells[:]); winner != models.Empty {
		targetBoard.Winner = winner
	} else if models.IsBoardFull(targetBoard.Cells[:]) {
		targetBoard.Winner = models.Tie
	}

	// Check Global Winner
	var globalCells [9]models.CellState
	for i := 0; i < 9; i++ {
		globalCells[i] = game.SubBoards[i].Winner
	}
	if globalWinner := models.CalculateWinner(globalCells[:]); globalWinner != models.Empty {
		game.Winner = globalWinner
		game.IsGameOver = true
	} else if models.IsBoardFull(globalCells[:]) {
		game.Winner = models.Tie
		game.IsGameOver = true
	}

	// Update Next Move Constraints
	game.NextBoardIdx = cellIdx
	if game.SubBoards[game.NextBoardIdx].Winner != models.Empty {
		game.NextBoardIdx = -1 // Target board is already won, player can go anywhere
	}

	// Switch Player (only if game not over)
	if !game.IsGameOver {
		if game.CurrentPlayer == models.PlayerX {
			game.CurrentPlayer = models.PlayerO
		} else {
			game.CurrentPlayer = models.PlayerX
		}
	}
}

// RequestBotMove asks the AI for a move and executes it.
func (s *gameService) RequestBotMove(ctx context.Context, id string) (*models.Game, error) {
	return nil, nil
}
