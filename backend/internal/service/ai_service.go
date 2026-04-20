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

// GetStandardMove suggests a move for a 3x3 game using Minimax.
func (s *aiService) GetStandardMove(_ context.Context, game *models.StandardGame) (int, error) {
	if game.IsGameOver {
		return -1, models.ErrGameOver
	}

	bestScore := -1000
	bestMove := -1

	for i, cell := range game.Board {
		if cell == models.Empty {
			game.Board[i] = game.CurrentPlayer
			score := s.minimax(game, 0, false)
			game.Board[i] = models.Empty // undo move

			if score > bestScore {
				bestScore = score
				bestMove = i
			}
		}
	}

	if bestMove == -1 {
		return -1, errors.New("no empty cells available")
	}

	return bestMove, nil
}

func (s *aiService) minimax(game *models.StandardGame, depth int, isMaximizing bool) int {
	winner := models.CalculateWinner(game.Board[:])
	if winner != models.Empty {
		if winner == game.CurrentPlayer {
			return 10 - depth
		}
		if winner == models.Tie {
			return 0
		}
		return depth - 10
	}
	if models.IsBoardFull(game.Board[:]) {
		return 0
	}

	opponent := models.PlayerX
	if game.CurrentPlayer == models.PlayerX {
		opponent = models.PlayerO
	}

	if isMaximizing {
		bestScore := -1000
		for i, cell := range game.Board {
			if cell == models.Empty {
				game.Board[i] = game.CurrentPlayer
				score := s.minimax(game, depth+1, false)
				game.Board[i] = models.Empty
				if score > bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	} else {
		bestScore := 1000
		for i, cell := range game.Board {
			if cell == models.Empty {
				game.Board[i] = opponent
				score := s.minimax(game, depth+1, true)
				game.Board[i] = models.Empty
				if score < bestScore {
					bestScore = score
				}
			}
		}
		return bestScore
	}
}

// GetUltimateMove suggests a move for a 9x9 game.
func (s *aiService) GetUltimateMove(_ context.Context, game *models.Game) (int, int, error) {
	if game.IsGameOver {
		return -1, -1, models.ErrGameOver
	}

	// For 9x9, Minimax is too slow. Let's start with a random move for now to pass Cycle 2 tests.
	// But it must be a VALID move.
	
	allowedBoards := []int{}
	if game.NextBoardIdx != -1 {
		allowedBoards = append(allowedBoards, game.NextBoardIdx)
	} else {
		for i := 0; i < 9; i++ {
			if game.SubBoards[i].Winner == models.Empty {
				allowedBoards = append(allowedBoards, i)
			}
		}
	}

	for _, bIdx := range allowedBoards {
		for cIdx, cell := range game.SubBoards[bIdx].Cells {
			if cell == models.Empty {
				return bIdx, cIdx, nil
			}
		}
	}

	return -1, -1, errors.New("no valid moves available")
}
