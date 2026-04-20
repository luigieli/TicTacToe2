package dto

import "tictactoe/internal/domain/models"

// MoveRequest defines the expected payload for making a move in Tic-Tac-Toe-2.
type MoveRequest struct {
	GameID   string `json:"game_id"`
	BoardIdx int    `json:"board_idx"` // 0-8
	CellIdx  int    `json:"cell_idx"`  // 0-8
}

// StandardMoveRequest defines the expected payload for making a move in Standard Tic-Tac-Toe.
type StandardMoveRequest struct {
	GameID  string `json:"game_id"`
	CellIdx int    `json:"cell_idx"` // 0-8
}

// CreateGameRequest defines the input for starting a new game.
type CreateGameRequest struct {
	Mode models.GameMode `json:"mode"` // PVP or PVE
}

// CreateGameResponse is returned when a new game is started.
type CreateGameResponse struct {
	GameID string `json:"game_id"`
}

// GameStateResponse represents the full game state sent to the client.
type GameStateResponse struct {
	ID            string           `json:"id"`
	Mode          models.GameMode  `json:"mode"`
	SubBoards     [9]models.Board  `json:"sub_boards"`
	CurrentPlayer models.CellState `json:"current_player"`
	NextBoardIdx  int              `json:"next_board_idx"`
	Winner        models.CellState `json:"winner"`
	IsGameOver    bool             `json:"is_game_over"`
}
