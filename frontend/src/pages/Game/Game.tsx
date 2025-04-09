import styles from "./Game.module.css";
import GameBoard from "../../components/GameBoard/GameBoard";
import Title from "../../components/Title/Title";
import LayoutBorder from "../../components/LayoutBorder/LayoutBorder";
import { useState } from "react";

type Player = "X" | "O" | null;
type Board = Player[];

function calculateWinner(board: (string | null)[]): string | null {
  const lines = [
    [0, 1, 2], [3, 4, 5], [6, 7, 8], // linhas
    [0, 3, 6], [1, 4, 7], [2, 5, 8], // colunas
    [0, 4, 8], [2, 4, 6],            // diagonais
  ];

  for (const [a, b, c] of lines) {
    if (board[a] && board[a] === board[b] && board[a] === board[c]) {
      return board[a];
    }
  }

  return null;
}

const Game: React.FC = () => {
  const [board, setBoard] = useState<Board>(Array(9).fill(null));
  const [currentPlayer, setCurrentPLayer] = useState<"X" | "O">("X");
  const [status, setStatus] = useState<string>("")
  const [winner, setWinner] = useState<string | null>("")
  
  const handleCellClick = (index: number) => {
    if (board[index] || calculateWinner(board)){
      setWinner(calculateWinner(board))
      setStatus(winner ? `Vencedor: ${winner}` : `Turno de: ${currentPlayer}`);
    };

    const newBoard = [...board];
    newBoard[index] = currentPlayer;
    setBoard(newBoard);
    setCurrentPLayer(currentPlayer === "X" ? "O" : "X");
  };

  // 

  return (
    <div className={styles.main}>
      <GameBoard squares={board} onSquareClick={handleCellClick}/>
      <span>{status}</span>
      <Title />
      <LayoutBorder />
    </div>
  );
};

export default Game;
