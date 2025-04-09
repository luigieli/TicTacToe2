import styles from "./GameBoard.module.css";
import Square from "../Square/Square";
import ScoreBoard from "../ScoreBoard/ScoreBoard";

type BoardProps = {
  squares: (string | null)[];
  onSquareClick: (index: number) => void;
};

const GameBoard: React.FC<BoardProps> = ({squares, onSquareClick}) => {
  return (
    <div className={styles.mainBoard}>
      <ScoreBoard />
      <div className={styles.board}>
        {squares.map((square: (string|null), i: number) =>(
            <Square key={i} value={square} onClick={()=>onSquareClick(i)}/>
        ))}
      </div>
    </div>
  );
};

export default GameBoard;
