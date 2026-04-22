package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"tictactoe/internal/domain/dto"
	"tictactoe/internal/domain/models"
	"tictactoe/internal/ports"
)

type StandardGameHandler struct {
	service ports.StandardGameService
}

// NewStandardGameHandler creates a new instance of StandardGameHandler.
func NewStandardGameHandler(service ports.StandardGameService) *StandardGameHandler {
	return &StandardGameHandler{service: service}
}

// CreateGame handles the POST / request.
func (h *StandardGameHandler) CreateGame(w http.ResponseWriter, r *http.Request) {
	var req dto.StandardCreateGameRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Printf("CreateGame: No valid body found, defaulting to PVP. Error: %v", err)
		req.Mode = string(models.PVP)
	}

	log.Printf("CreateGame: Initializing game with mode: %s", req.Mode)
	gameID, err := h.service.CreateGame(r.Context(), req.Mode)
	if err != nil {
		h.respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	h.respondWithJSON(w, http.StatusCreated, dto.CreateGameResponse{GameID: gameID})
}

// GetGameState handles the GET /{id} request.
func (h *StandardGameHandler) GetGameState(w http.ResponseWriter, r *http.Request) {
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

// MakeMove handles the POST /{id}/move request.
func (h *StandardGameHandler) MakeMove(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		h.respondWithError(w, http.StatusBadRequest, "game id is required")
		return
	}

	var req dto.StandardMoveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.respondWithError(w, http.StatusBadRequest, "invalid request body")
		return
	}

	log.Printf("MakeMove: id=%s, cell=%d", id, req.CellIdx)
	game, err := h.service.MakeMove(r.Context(), id, req.CellIdx)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, models.ErrInvalidMove) || errors.Is(err, models.ErrGameOver) ||
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

func (h *StandardGameHandler) respondWithError(w http.ResponseWriter, code int, message string) {
	h.respondWithJSON(w, code, map[string]string{"error": message})
}

func (h *StandardGameHandler) respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
