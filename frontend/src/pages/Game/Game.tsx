import styles from "./Game.module.css";
import GameBoard from "../../components/GameBoard/GameBoard";
import Title from "../../components/Title/Title";
import LayoutBorder from "../../components/FrameBorder/FrameBorder";
import calculateWinner from "../../utils/CalculateWinner";
import { useState } from "react";
import { PLAYERS, COLORS } from "../../constants";

type Player = "X" | "O" | null;
type Board = Player[];

const Game: React.FC = () => {
  const [board, setBoard] = useState<Board>(Array(9).fill(null));
  const [currentPlayer, setCurrentPLayer] = useState<"X" | "O">(PLAYERS.X);
  const [scores, setScore] = useState({
    playerX: 0,
    ties: 0,
    playerO: 0,
  });

  // Mapeia as cores com base no valor de cada célula do tabuleiro
  const boardColors = board.map((square) =>
    square === "O" ? COLORS.PLAYER_O : square === "X" ? COLORS.PLAYER_X : ""
  );

  const handleCellClick = (index: number) => {
    if (board[index] || calculateWinner(board)) return;

    const newBoard = [...board];
    newBoard[index] = currentPlayer;
    setBoard(newBoard);
    setCurrentPLayer(currentPlayer === PLAYERS.X ? PLAYERS.O : PLAYERS.X);
  };

  return (
    <div className={styles.main}>
      <GameBoard
        scores={scores}
        squares={board}
        colors={boardColors} // Passa as cores como prop
        onSquareClick={handleCellClick}
      />
      <Title />
      <LayoutBorder />
    </div>
  );
};

export default Game;
