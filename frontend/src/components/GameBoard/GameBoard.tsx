import styles from "./GameBoard.module.css";
import Square from "../Square/Square";
import ScoreBoard from "../ScoreBoard/ScoreBoard";

type BoardProps = {
  squares: (string | null)[];
  onSquareClick: (index: number) => void;
  scores: {
    playerX: number;
    ties: number;
    playerO: number;
  };
  winner: string;
  isGameOver: boolean;
  isPVE?: boolean;
};

const GameBoard: React.FC<BoardProps> = ({
  squares,
  onSquareClick,
  scores,
  winner,
  isGameOver,
  isPVE,
}) => {
  const getWinnerText = () => {
    if (winner === "X") return "JOGADOR X VENCEU!";
    if (winner === "O") return isPVE ? "A CPU VENCEU!" : "JOGADOR O VENCEU!";
    if (winner === "Z") return "EMPATE!";
    return "";
  };

  return (
    <div className={styles.mainBoard}>
      <ScoreBoard scores={scores} isPVE={isPVE} />
      
      <div className={`${styles.boardContainer} ${isGameOver ? styles.gameOver : ''}`}>
        <div className={styles.board}>
          {squares.map((square: string | null, i: number) => (
            <Square
              key={i}
              value={square}
              onClick={() => onSquareClick(i)}
              disabled={isGameOver}
            />
          ))}
        </div>

        {isGameOver && (
          <div className={styles.overlay}>
            <div className={styles.winnerCard}>
              <span className={styles.winnerLabel}>FIM DE JOGO</span>
              <h2 className={`${styles.winnerTitle} ${winner === 'X' ? styles.xColor : winner === 'O' ? styles.oColor : styles.tieColor}`}>
                {getWinnerText()}
              </h2>
            </div>
          </div>
        )}
      </div>
    </div>
  );
};

export default GameBoard;
