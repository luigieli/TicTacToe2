package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"tictactoe/internal/domain/dto"
	"tictactoe/internal/domain/models"
	"tictactoe/internal/ports"
)

type TicTacToe2Handler struct {
	service ports.GameService
}

// NewTicTacToe2Handler creates a new instance of TicTacToe2Handler.
func NewTicTacToe2Handler(service ports.GameService) *TicTacToe2Handler {
	return &TicTacToe2Handler{service: service}
}

// CreateGame handles the POST /game request.
func (h *TicTacToe2Handler) CreateGame(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		req.Mode = models.ModePVP
	}

	gameID, err := h.service.CreateGame(r.Context(), req)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusCreated, dto.CreateGameResponse{GameID: gameID})
}

// GetGameState handles the GET /game/{id} request.
func (h *TicTacToe2Handler) GetGameState(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "game id is required")
		return
	}

	game, err := h.service.GetGameState(r.Context(), id)
	if err != nil {
		if errors.Is(err, models.ErrGameNotFound) {
			h.respondWithError(w, http.StatusNotFound, err.Error())
			return
		}
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, game)
}

// MakeMove handles the POST /move request.
func (h *TicTacToe2Handler) MakeMove(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "game id is required")
		return
	}

	var req dto.MoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	req.GameID = id

	game, err := h.service.MakeMove(r.Context(), req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, models.ErrInvalidMove) || errors.Is(err, models.ErrWrongBoard) ||
			errors.Is(err, models.ErrBoardAlreadyWon) || errors.Is(err, models.ErrGameOver) ||
			errors.Is(err, models.ErrCellAlreadyTaken) {
			status = http.StatusBadRequest
		} else if errors.Is(err, models.ErrGameNotFound) {
			status = http.StatusNotFound
		}
		h.respondWithError(w, status, err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusOK, game)
}

func (h *TicTacToe2Handler) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJSON(w, code, map[string]string{"error": message})
}

func (h *TicTacToe2Handler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
