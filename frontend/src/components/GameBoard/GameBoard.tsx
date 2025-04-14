import styles from "./GameBoard.module.css";
import Square from "../Square/Square";
import ScoreBoard from "../ScoreBoard/ScoreBoard";
import NewGameButton from "../NewGameButton/NewGameButton";

type BoardProps = {
  squares: ("X" | "O" | null)[];
  colors: string[]; // Nova prop para as cores
  onSquareClick: (index: number) => void;
  scores: {
    playerX: number;
    ties: number;
    playerO: number;
  };
};

const GameBoard: React.FC<BoardProps> = ({
  squares,
  colors,
  onSquareClick,
  scores,
}) => {
  return (
    <div className={styles.mainBoard}>
      <NewGameButton />
      <ScoreBoard scores={scores} />
      <div className={styles.board}>
        {squares.map((square: "X" | "O" | null, i: number) => (
          <Square
            key={i}
            value={square}
            onClick={() => onSquareClick(i)}
            color={colors[i]} // Usa a cor correspondente
          />
        ))}
      </div>
    </div>
  );
};

export default GameBoard;
