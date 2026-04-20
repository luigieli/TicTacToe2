const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export type GameMode = "PVP" | "PVE";

export interface StandardGame {
  id: string;
  mode: GameMode;
  board: ("X" | "O" | "")[];
  current_player: "X" | "O";
  winner: "X" | "O" | "Z" | ""; // "Z" is Tie
  is_game_over: boolean;
}

export interface SubBoard {
  cells: ("X" | "O" | "")[];
  winner: "X" | "O" | "Z" | "";
}

export interface UltimateGame {
  id: string;
  mode: GameMode;
  sub_boards: SubBoard[];
  current_player: "X" | "O";
  next_board_idx: number; // -1 for any
  winner: "X" | "O" | "Z" | "";
  is_game_over: boolean;
}

/**
 * STANDARD (3x3) API
 */
export async function createStandardGame(mode: GameMode): Promise<string> {
  const response = await fetch(`${API_BASE_URL}/games/standard`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ mode }),
  });
  if (!response.ok) throw new Error("Failed to create standard game");
  const data = await response.json();
  return data.game_id;
}

export async function getStandardGameState(id: string): Promise<StandardGame> {
  const response = await fetch(`${API_BASE_URL}/games/standard/${id}`);
  if (!response.ok) throw new Error("Failed to fetch standard game state");
  return response.json();
}

export async function makeStandardMove(id: string, cellIdx: number): Promise<StandardGame> {
  const response = await fetch(`${API_BASE_URL}/games/standard/${id}/move`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ cell_idx: cellIdx }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.error || "Failed to make move");
  }
  return response.json();
}

/**
 * ULTIMATE (9x9) API
 */
export async function createUltimateGame(mode: GameMode): Promise<string> {
  const response = await fetch(`${API_BASE_URL}/games/tictactoe2`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ mode }),
  });
  if (!response.ok) throw new Error("Failed to create ultimate game");
  const data = await response.json();
  return data.game_id;
}

export async function getUltimateGameState(id: string): Promise<UltimateGame> {
  const response = await fetch(`${API_BASE_URL}/games/tictactoe2/${id}`);
  if (!response.ok) throw new Error("Failed to fetch ultimate game state");
  return response.json();
}

export async function makeUltimateMove(id: string, boardIdx: number, cellIdx: number): Promise<UltimateGame> {
  const response = await fetch(`${API_BASE_URL}/games/tictactoe2/${id}/move`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ board_idx: boardIdx, cell_idx: cellIdx }),
  });
  if (!response.ok) {
    const errorData = await response.json();
    throw new Error(errorData.error || "Failed to make move");
  }
  return response.json();
}
