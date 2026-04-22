const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export type GameMode = "PVP" | "PVE_EASY" | "PVE_MEDIUM" | "PVE_HARD";

export interface StandardGame {
  id: string;
  mode: GameMode;
  board: ("X" | "O" | "")[];
  current_player: "X" | "O";
  winner: "X" | "O" | "Z" | ""; // "Z" is Tie
  is_game_over: boolean;
}

export interface Game {
  id: string;
  current_player: "X" | "O";
  next_board_idx: number;
  winner: "X" | "O" | "Z" | "";
  is_game_over: boolean;
  sub_boards: {
    cells: ("X" | "O" | "")[];
    winner: "X" | "O" | "Z" | "";
  }[];
}

/**
 * STANDARD (3x3) API
 */
export async function createGame(mode: GameMode = "PVP"): Promise<string> {
  const response = await fetch(`${API_BASE_URL}/games/standard/`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ mode }),
  });

  if (!response.ok) {
    throw new Error("Failed to create game");
  }

  const data = await response.json();
  return data.game_id;
}

/**
 * Retrieves the current state of a game.
 */
export async function getGameState(id: string): Promise<StandardGame> {
  const response = await fetch(`${API_BASE_URL}/games/standard/${id}`);

  if (!response.ok) {
    throw new Error("Failed to fetch game state");
  }

  return response.json();
}

/**
 * Makes a move in the specified game.
 */
export async function makeMove(id: string, cellIdx: number): Promise<StandardGame> {
  const response = await fetch(`${API_BASE_URL}/games/standard/${id}/move`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ cell_idx: cellIdx }),
  });

  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.error || "Failed to make move");
  }

  return response.json();
}
