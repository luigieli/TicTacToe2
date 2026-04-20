package models

import "github.com/google/uuid"

// StandardGame represents the state of a traditional 3x3 Tic-Tac-Toe game.
type StandardGame struct {
	ID            string       `json:"id"`
	Mode          GameMode     `json:"mode"`
	Board         [9]CellState `json:"board"`
	CurrentPlayer CellState    `json:"current_player"`
	Winner        CellState    `json:"winner"` // Empty, PlayerX, PlayerO, or Tie
	IsGameOver    bool         `json:"is_game_over"`
}

// NewStandardGame initializes a new 3x3 game.
func NewStandardGame(mode GameMode) *StandardGame {
	if mode == "" {
		mode = ModePVP
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
