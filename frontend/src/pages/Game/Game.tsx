import styles from "./Game.module.css";
import GameBoard from "../../components/GameBoard/GameBoard";
import { useEffect, useState, useCallback } from "react";
import { createGame, makeMove, StandardGame, GameMode } from "../../utils/api";

interface GameProps {
  mode: GameMode;
  onBack: () => void;
}

const Game: React.FC<GameProps> = ({ mode, onBack }) => {
  const [gameID, setGameID] = useState<string | null>(null);
  const [gameState, setGameState] = useState<StandardGame | null>(null);
  const [errorMessage, setErrorMessage] = useState<string | null>(null);
  const [scores, setScore] = useState({
    playerX: 0,
    ties: 0,
    playerO: 0,
  });

  const initGame = useCallback(async () => {
    try {
      const id = await createGame(mode);
      setGameID(id);
      setGameState({
        id: id,
        mode: mode,
        board: Array(9).fill(""),
        current_player: "X",
        winner: "",
        is_game_over: false,
      });
      setErrorMessage(null);
    } catch (err: any) {
      setErrorMessage("Could not initialize game: " + err.message);
    }
  }, [mode]);

  useEffect(() => {
    initGame();
  }, [initGame]);

  const handleCellClick = async (index: number) => {
    if (!gameID || !gameState) return;
    if (gameState.is_game_over || gameState.board[index] !== "") return;

    try {
      const updatedState = await makeMove(gameID, index);
      setGameState(updatedState);
      setErrorMessage(null);

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
    } else if (winner === "Z") {
      setScore((prev) => ({ ...prev, ties: prev.ties + 1 }));
    }
  };

  const handleRestart = () => {
    initGame();
  };

  const boardCells = gameState?.board.map(cell => {
    if (cell === "X") return "X";
    if (cell === "O") return "O";
    return null;
  }) || Array(9).fill(null);

  return (
    <div className={styles.container}>
      <header className={styles.header}>
        <button className={styles.backButton} onClick={onBack}>
          <span>←</span> Menu
        </button>
      </header>

      <main className={styles.gameCard}>
        <GameBoard
          scores={scores}
          squares={boardCells}
          onSquareClick={handleCellClick}
          winner={gameState?.winner || ""}
          isGameOver={gameState?.is_game_over || false}
          isPVE={mode.startsWith("PVE")}
        />
        
        <div className={styles.controls}>
          <button className={styles.actionButton} onClick={handleRestart}>
            Reiniciar Partida
          </button>
        </div>
      </main>

      {errorMessage && (
        <div className={styles.error}>
          {errorMessage}
        </div>
      )}
    </div>
  );
};

export default Game;
