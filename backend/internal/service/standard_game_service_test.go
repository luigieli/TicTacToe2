package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"tictactoe/internal/domain/dto"
	"tictactoe/internal/domain/models"
	"tictactoe/internal/repository"
)

func TestStandardGameService(t *testing.T) {
	repo := repository.NewStandardGameRepository()
	ai := NewAIService()
	svc := NewStandardGameService(repo, ai)
	ctx := context.Background()

	t.Run("CreateGame", func(t *testing.T) {
		id, err := svc.CreateGame(ctx, dto.CreateGameRequest{Mode: models.ModePVP})
		assert.NoError(t, err)
		assert.NotEmpty(t, id)

		game, err := svc.GetGameState(ctx, id)
		assert.NoError(t, err)
		assert.Equal(t, id, game.ID)
		assert.Equal(t, models.ModePVP, game.Mode)
		assert.Equal(t, models.PlayerX, game.CurrentPlayer)
		assert.False(t, game.IsGameOver)
		for _, cell := range game.Board {
			assert.Equal(t, models.Empty, cell)
		}
	})

	t.Run("MakeMove_PVP", func(t *testing.T) {
		id, _ := svc.CreateGame(ctx, dto.CreateGameRequest{Mode: models.ModePVP})

		// X moves to center
		game, err := svc.MakeMove(ctx, id, 4)
		assert.NoError(t, err)
		assert.Equal(t, models.PlayerX, game.Board[4])
		assert.Equal(t, models.PlayerO, game.CurrentPlayer)

		// O moves to top-left
		game, err = svc.MakeMove(ctx, id, 0)
		assert.NoError(t, err)
		assert.Equal(t, models.PlayerO, game.Board[0])
		assert.Equal(t, models.PlayerX, game.CurrentPlayer)
	})

	t.Run("MakeMove_PVE_AutoResponse", func(t *testing.T) {
		id, _ := svc.CreateGame(ctx, dto.CreateGameRequest{Mode: models.ModePVE})
		
		// X moves to center. AI should respond immediately.
		game, err := svc.MakeMove(ctx, id, 4)
		assert.NoError(t, err)
		
		// X is at 4, O (AI) should have played elsewhere.
		assert.Equal(t, models.PlayerX, game.Board[4])
		
		filledCount := 0
		for _, cell := range game.Board {
			if cell != models.Empty {
				filledCount++
			}
		}
		// Should be 2 (X and the AI's response O)
		assert.Equal(t, 2, filledCount)
		assert.Equal(t, models.PlayerX, game.CurrentPlayer) // Back to X's turn
	})

	t.Run("InvalidMoves", func(t *testing.T) {
		id, _ := svc.CreateGame(ctx, dto.CreateGameRequest{Mode: models.ModePVP})
		svc.MakeMove(ctx, id, 4) // X moves to 4

		// O tries to move to 4 (already taken)
		_, err := svc.MakeMove(ctx, id, 4)
		assert.Error(t, err)
		assert.ErrorIs(t, err, models.ErrCellAlreadyTaken)

		// O tries to move out of bounds
		_, err = svc.MakeMove(ctx, id, 9)
		assert.Error(t, err)
		assert.ErrorIs(t, err, models.ErrInvalidMove)
	})

	t.Run("WinCondition", func(t *testing.T) {
		id, _ := svc.CreateGame(ctx, dto.CreateGameRequest{Mode: models.ModePVP})
		// X: 0, 1, 2
		// O: 3, 4
		svc.MakeMove(ctx, id, 0)              // X
		svc.MakeMove(ctx, id, 3)              // O
		svc.MakeMove(ctx, id, 1)              // X
		svc.MakeMove(ctx, id, 4)              // O
		game, err := svc.MakeMove(ctx, id, 2) // X wins

		assert.NoError(t, err)
		assert.True(t, game.IsGameOver)
		assert.Equal(t, models.PlayerX, game.Winner)
	})
}
