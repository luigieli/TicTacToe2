import styles from "./Game.module.css";
import GameBoard from "../../components/GameBoard/GameBoard";
import Title from "../../components/Title/Title";
import LayoutBorder from "../../components/LayoutBorder/LayoutBorder";
import calculateWinner from "../../utils/CalculateWinner";
import { useState } from "react";

type Player = "X" | "O" | null;
type Board = Player[];

const Game: React.FC = () => {
  const [board, setBoard] = useState<Board>(Array(9).fill(null));
  const [currentPlayer, setCurrentPLayer] = useState<"X" | "O">("X");
  const [scores, setScore] = useState({
    playerX: 0,
    ties: 0,
    playerO: 0,
  })

  const handleCellClick = (index: number) => {
    if (board[index] || calculateWinner(board)) return;

    const newBoard = [...board];
    newBoard[index] = currentPlayer;
    setBoard(newBoard);
    setCurrentPLayer(currentPlayer === "X" ? "O" : "X");
  };
  
  return (
    <div className={styles.main}>
      <GameBoard
        scores={scores}
        squares={board}
        onSquareClick={handleCellClick}
      />
      <Title />
      <LayoutBorder />
    </div>
  );
};

export default Game;
