const API_BASE_URL = import.meta.env.VITE_API_BASE_URL;

export interface StandardGame {
  id: string;
  board: ("X" | "O" | "")[];
  current_player: "X" | "O";
  winner: "X" | "O" | "Z" | ""; // "Z" is Tie
  is_game_over: boolean;
}

/**
 * Creates a new standard Tic-Tac-Toe game.
 */
export async function createGame(): Promise<string> {
  const response = await fetch(`${API_BASE_URL}/`, {
    method: "POST",
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
  const response = await fetch(`${API_BASE_URL}/${id}`);

  if (!response.ok) {
    throw new Error("Failed to fetch game state");
  }

  return response.json();
}

/**
 * Makes a move in the specified game.
 */
export async function makeMove(id: string, cellIdx: number): Promise<StandardGame> {
  const response = await fetch(`${API_BASE_URL}/${id}/move`, {
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
