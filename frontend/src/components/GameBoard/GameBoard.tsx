import styles from "./GameBoard.module.css";
import Square from "../Square/Square";
import ScoreBoard from "../ScoreBoard/ScoreBoard";
import NewGameButton from "../NewGameButton/NewGameButton";

type BoardProps = {
  squares: (string | null)[];
  onSquareClick: (index: number) => void;
  scores: {
    playerX: number;
    ties: number;
    playerO: number;
  };
};

function textPlayerColor(player: string | null) {
  if (player == null) return "";
  return player === "O" ? "#E2BE00" : "#72CFF9";
}

const GameBoard: React.FC<BoardProps> = ({
  squares,
  onSquareClick,
  scores,
}) => {
  return (
    <div className={styles.mainBoard}>
      <NewGameButton />
      <ScoreBoard scores={scores} />
      <div className={styles.board}>
        {squares.map((square: string | null, i: number) => (
          <Square
            key={i}
            value={square}
            onClick={() => onSquareClick(i)}
            color={textPlayerColor(square)}
          />
        ))}
      </div>
    </div>
  );
};

export default GameBoard;
