package models

import "errors"

// CellState represents the state of a single cell in a 3x3 board.
type CellState string

const (
	Empty   CellState = ""
	PlayerX CellState = "X"
	PlayerO CellState = "O"
	Tie     CellState = "Z"
)

// Board represents a single 3x3 grid.
type Board struct {
	Cells  [9]CellState `json:"cells"`
	Winner CellState    `json:"winner"`
}

// GameMode represents the type of opponents in a game.
type GameMode string

const (
	ModePVP GameMode = "PVP"
	ModePVE GameMode = "PVE"
)

// Game represents the state of an Ultimate Tic Tac Toe game.
type Game struct {
	ID            string    `json:"id"`
	Mode          GameMode  `json:"mode"`
	SubBoards     [9]Board  `json:"sub_boards"`
	CurrentPlayer CellState `json:"current_player"`
	NextBoardIdx  int       `json:"next_board_idx"` // -1 means any board is allowed
	Winner        CellState `json:"winner"`
	IsGameOver    bool      `json:"is_game_over"`
}

// Errors
var (
	ErrGameNotFound     = errors.New("game not found")
	ErrInvalidMove      = errors.New("invalid move")
	ErrBoardAlreadyWon  = errors.New("board already won")
	ErrWrongBoard       = errors.New("must play in the mandated sub-board")
	ErrGameOver         = errors.New("game is already over")
	ErrCellAlreadyTaken = errors.New("cell is already occupied")
)
