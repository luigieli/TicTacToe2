import styles from "./Game.module.css";
import GameBoard from "../../components/GameBoard/GameBoard";
import Title from "../../components/Title/Title";
import LayoutBorder from "../../components/FrameBorder/FrameBorder";
import { useEffect, useState } from "react";
import { createStandardGame, makeStandardMove, StandardGame } from "../../utils/api";

const Game: React.FC = () => {
  const [gameID, setGameID] = useState<string | null>(null);
  const [gameState, setGameState] = useState<StandardGame | null>(null);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [scores, setScore] = useState({
    playerX: 0,
    ties: 0,
    playerO: 0,
  });

  // Initialize game session on mount.
  useEffect(() => {
    async function init() {
      try {
        const id = await createStandardGame("PVP"); // Defaulting to PVP until Menu is ready
        setGameID(id);
        // Initial state
        setGameState({
          id: id,
          mode: "PVP",
          board: Array(9).fill(""),
          current_player: "X",
          winner: "",
          is_game_over: false,
        });
      } catch (err: any) {
        setErrorMessage("Could not initialize game: " + err.message);
      }
    }
    init();
  }, []);

  const handleCellClick = async (index: number) => {
    if (!gameID || !gameState) return;
    if (gameState.is_game_over || gameState.board[index] !== "") return;

    try {
      const updatedState = await makeStandardMove(gameID, index);
      setGameState(updatedState);
      setErrorMessage(null); // Clear errors on success

      if (updatedState.is_game_over) {
        updateScores(updatedState.winner);
      }
    } catch (err: any) {
      setErrorMessage(err.message);
    }
  };

  const updateScores = (winner: string) => {
    if (winner === "X") {
      setScore((prev) => ({ ...prev, playerX: prev.playerX + 1 }));
    } else if (winner === "O") {
      setScore((prev) => ({ ...prev, playerO: prev.playerO + 1 }));
    } else if (winner === "Z") { // backend Tie is "Z"
      setScore((prev) => ({ ...prev, ties: prev.ties + 1 }));
    }
  };

  // Convert string[] from backend to Player[] for GameBoard
  const boardCells = gameState?.board.map(cell => {
    if (cell === "X") return "X";
    if (cell === "O") return "O";
    return null;
  }) || Array(9).fill(null);

  return (
    <div className={styles.main}>
      <GameBoard
        scores={scores}
        squares={boardCells}
        onSquareClick={handleCellClick}
      />
      <Title />
      <LayoutBorder />
      {errorMessage && (
        <div style={{ color: "red", position: "absolute", bottom: "20px" }}>
          {errorMessage}
        </div>
      )}
    </div>
  );
};

export default Game;
