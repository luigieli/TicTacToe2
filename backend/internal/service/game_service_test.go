package service

import (
	"context"
	"testing"

	"tictactoe/internal/domain/dto"
	"tictactoe/internal/domain/models"
)

// Mock repository for testing
type mockRepo struct {
	games map[string]*models.Game
}

func (m *mockRepo) Save(ctx context.Context, game *models.Game) error {
	m.games[game.ID] = game
	return nil
}

func (m *mockRepo) FindByID(ctx context.Context, id string) (*models.Game, error) {
	game, ok := m.games[id]
	if !ok {
		return nil, models.ErrGameNotFound
	}
	return game, nil
}

func newMockRepo() *mockRepo {
	return &mockRepo{games: make(map[string]*models.Game)}
}

func TestCreateGame(t *testing.T) {
	repo := newMockRepo()
	ai := NewAIService()
	svc := NewGameService(repo, ai)

	gameID, err := svc.CreateGame(context.Background(), dto.CreateGameRequest{Mode: models.ModePVP})
	if err != nil {
		t.Fatalf("failed to create game: %v", err)
	}

	if gameID == "" {
		t.Error("expected non-empty game ID")
	}

	game, err := repo.FindByID(context.Background(), gameID)
	if err != nil {
		t.Fatalf("failed to find created game in repo: %v", err)
	}

	if game.Mode != models.ModePVP {
		t.Errorf("expected Mode PVP, got %v", game.Mode)
	}

	if game.CurrentPlayer != models.PlayerX {
		t.Errorf("expected starting player X, got %v", game.CurrentPlayer)
	}
}

func TestMakeMove(t *testing.T) {
	repo := newMockRepo()
	ai := NewAIService()
	svc := NewGameService(repo, ai)

	gameID, _ := svc.CreateGame(context.Background(), dto.CreateGameRequest{Mode: models.ModePVP})

	// Move Player X to Board 0, Cell 4
	req := dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 0,
		CellIdx:  4,
	}

	game, err := svc.MakeMove(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to make move: %v", err)
	}

	if game.SubBoards[0].Cells[4] != models.PlayerX {
		t.Errorf("expected Board 0, Cell 4 to be PlayerX, got %v", game.SubBoards[0].Cells[4])
	}

	if game.CurrentPlayer != models.PlayerO {
		t.Errorf("expected current player to be O, got %v", game.CurrentPlayer)
	}
}

func TestMakeMove_PVE_AutoResponse(t *testing.T) {
	repo := newMockRepo()
	ai := NewAIService()
	svc := NewGameService(repo, ai)

	gameID, _ := svc.CreateGame(context.Background(), dto.CreateGameRequest{Mode: models.ModePVE})

	// X moves to Board 0, Cell 4. Bot should respond.
	req := dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 0,
		CellIdx:  4,
	}

	game, err := svc.MakeMove(context.Background(), req)
	if err != nil {
		t.Fatalf("failed to make move: %v", err)
	}

	// Two moves should be present
	filledCount := 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if game.SubBoards[i].Cells[j] != models.Empty {
				filledCount++
			}
		}
	}

	if filledCount != 2 {
		t.Errorf("expected 2 filled cells, got %d", filledCount)
	}

	if game.CurrentPlayer != models.PlayerX {
		t.Errorf("expected back to player X, got %v", game.CurrentPlayer)
	}
}
