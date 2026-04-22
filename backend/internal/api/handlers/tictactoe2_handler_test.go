package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"tictactoe/internal/domain/dto"
	"tictactoe/internal/domain/models"
)

// Mock service for testing
type mockService struct {
	game *models.Game
}

func (m *mockService) CreateGame(ctx context.Context, req dto.CreateGameRequest) (string, error) {
	return "test-id", nil
}

func (m *mockService) GetGameState(ctx context.Context, id string) (*models.Game, error) {
	if id == "test-id" {
		return m.game, nil
	}
	return nil, models.ErrGameNotFound
}

func (m *mockService) MakeMove(ctx context.Context, req dto.MoveRequest) (*models.Game, error) {
	if req.GameID == "test-id" {
		m.game.SubBoards[req.BoardIdx].Cells[req.CellIdx] = models.PlayerX
		return m.game, nil
	}
	return nil, models.ErrGameNotFound
}

func setupTestRouter(h *TicTacToe2Handler) *chi.Mux {
	r := chi.NewRouter()
	r.Route("/games", func(r chi.Router) {
		r.Post("/", h.CreateGame)
		r.Get("/{id}", h.GetGameState)
		r.Post("/{id}/move", h.MakeMove)
	})
	return r
}

func TestGetGameState_REST(t *testing.T) {
	svc := &mockService{game: &models.Game{ID: "test-id"}}
	handler := NewTicTacToe2Handler(svc)
	router := setupTestRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/games/test-id", nil)
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d. Body: %s", rr.Code, rr.Body.String())
	}

	var game models.Game
	if err := json.NewDecoder(rr.Body).Decode(&game); err != nil {
		t.Fatal(err)
	}

	if game.ID != "test-id" {
		t.Errorf("expected game ID test-id, got %s", game.ID)
	}
}

func TestMakeMove_REST(t *testing.T) {
	svc := &mockService{game: &models.Game{ID: "test-id"}}
	handler := NewTicTacToe2Handler(svc)
	router := setupTestRouter(handler)

	move := dto.MoveRequest{BoardIdx: 0, CellIdx: 4}
	body, _ := json.Marshal(move)

	// Note: GameID is now in the URL, not the body (or both)
	req := httptest.NewRequest(http.MethodPost, "/games/test-id/move", bytes.NewBuffer(body))
	rr := httptest.NewRecorder()

	router.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d. Body: %s", rr.Code, rr.Body.String())
	}
}
