package models

import "github.com/google/uuid"

// StandardGameMode defines if it's PVP or PVE with difficulties.
type StandardGameMode string

const (
	PVP        StandardGameMode = "PVP"
	PVE_EASY   StandardGameMode = "PVE_EASY"
	PVE_MEDIUM StandardGameMode = "PVE_MEDIUM"
	PVE_HARD   StandardGameMode = "PVE_HARD"
)

// StandardGame represents the state of a traditional 3x3 Tic-Tac-Toe game.
type StandardGame struct {
	ID            string           `json:"id"`
	Mode          StandardGameMode `json:"mode"`
	Board         [9]CellState     `json:"board"`
	CurrentPlayer CellState        `json:"current_player"`
	Winner        CellState        `json:"winner"` // Empty, PlayerX, PlayerO, or Tie
	IsGameOver    bool             `json:"is_game_over"`
}

// NewStandardGame initializes a new 3x3 game.
func NewStandardGame(mode StandardGameMode) *StandardGame {
	if mode == "" {
		mode = PVP
	}
	return &StandardGame{
		ID:            uuid.New().String(),
		Mode:          mode,
		Board:         [9]CellState{},
		CurrentPlayer: PlayerX,
		IsGameOver:    false,
		Winner:        Empty,
	}
}
