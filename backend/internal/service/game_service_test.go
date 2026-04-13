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

func newMockRepo() *mockRepo {
	return &mockRepo{games: make(map[string]*models.Game)}
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

func TestCreateGame(t *testing.T) {
	repo := newMockRepo()
	svc := NewGameService(repo)

	gameID, err := svc.CreateGame(context.Background())
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

	if game.CurrentPlayer != models.PlayerX {
		t.Errorf("expected starting player X, got %v", game.CurrentPlayer)
	}

	if game.NextBoardIdx != -1 {
		t.Errorf("expected first move to be anywhere (-1), got %v", game.NextBoardIdx)
	}
}

func TestMakeMove(t *testing.T) {
	repo := newMockRepo()
	svc := NewGameService(repo)

	gameID, _ := svc.CreateGame(context.Background())

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

	if game.NextBoardIdx != 4 {
		t.Errorf("expected next move to be restricted to Board 4, got %v", game.NextBoardIdx)
	}
}

func TestGetGameState(t *testing.T) {
	repo := newMockRepo()
	svc := NewGameService(repo)

	gameID, _ := svc.CreateGame(context.Background())

	game, err := svc.GetGameState(context.Background(), gameID)
	if err != nil {
		t.Fatalf("failed to get game state: %v", err)
	}

	if game.ID != gameID {
		t.Errorf("expected game ID %s, got %s", gameID, game.ID)
	}
}

func TestGetGameState_NotFound(t *testing.T) {
	repo := newMockRepo()
	svc := NewGameService(repo)

	_, err := svc.GetGameState(context.Background(), "nonexistent")
	if err != models.ErrGameNotFound {
		t.Errorf("expected ErrGameNotFound, got %v", err)
	}
}

func TestMakeMove_InvalidCellIndex(t *testing.T) {
	repo := newMockRepo()
	svc := NewGameService(repo)

	gameID, _ := svc.CreateGame(context.Background())

	req := dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 0,
		CellIdx:  10, // Invalid
	}

	_, err := svc.MakeMove(context.Background(), req)
	if err != models.ErrInvalidMove {
		t.Errorf("expected ErrInvalidMove, got %v", err)
	}
}

func TestMakeMove_InvalidBoardIndex(t *testing.T) {
	repo := newMockRepo()
	svc := NewGameService(repo)

	gameID, _ := svc.CreateGame(context.Background())

	req := dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 10, // Invalid
		CellIdx:  0,
	}

	_, err := svc.MakeMove(context.Background(), req)
	if err != models.ErrInvalidMove {
		t.Errorf("expected ErrInvalidMove, got %v", err)
	}
}

func TestMakeMove_CellAlreadyTaken(t *testing.T) {
	repo := newMockRepo()
	svc := NewGameService(repo)

	gameID, _ := svc.CreateGame(context.Background())

	// Player X moves at Board 0, Cell 4. Next move must be in Board 4.
	svc.MakeMove(context.Background(), dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 0,
		CellIdx:  4,
	})

	// Player O moves at Board 4, Cell 0. Next move must be in Board 0.
	svc.MakeMove(context.Background(), dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 4,
		CellIdx:  0,
	})

	// Player X tries to play same cell as the first move.
	req := dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 0,
		CellIdx:  4, // Already taken by Player X
	}

	_, err := svc.MakeMove(context.Background(), req)
	if err != models.ErrCellAlreadyTaken {
		t.Errorf("expected ErrCellAlreadyTaken, got %v", err)
	}
}

func TestMakeMove_WrongBoard(t *testing.T) {
	repo := newMockRepo()
	svc := NewGameService(repo)

	gameID, _ := svc.CreateGame(context.Background())

	// First move in board 0, cell 0 (sends next player to board 0)
	svc.MakeMove(context.Background(), dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 0,
		CellIdx:  0,
	})

	// Try to play in board 1 instead of board 0
	req := dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 1, // Wrong board!
		CellIdx:  0,
	}

	_, err := svc.MakeMove(context.Background(), req)
	if err != models.ErrWrongBoard {
		t.Errorf("expected ErrWrongBoard, got %v", err)
	}
}

func TestMakeMove_PlayerSwitch(t *testing.T) {
	repo := newMockRepo()
	svc := NewGameService(repo)

	gameID, _ := svc.CreateGame(context.Background())

	game, _ := svc.GetGameState(context.Background(), gameID)
	if game.CurrentPlayer != models.PlayerX {
		t.Errorf("expected starting player X, got %v", game.CurrentPlayer)
	}

	// X makes a move
	game, _ = svc.MakeMove(context.Background(), dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 0,
		CellIdx:  0,
	})

	if game.CurrentPlayer != models.PlayerO {
		t.Errorf("expected current player to be O, got %v", game.CurrentPlayer)
	}

	// O makes a move
	game, _ = svc.MakeMove(context.Background(), dto.MoveRequest{
		GameID:   gameID,
		BoardIdx: 0,
		CellIdx:  1,
	})

	if game.CurrentPlayer != models.PlayerX {
		t.Errorf("expected current player to be X, got %v", game.CurrentPlayer)
	}
}
