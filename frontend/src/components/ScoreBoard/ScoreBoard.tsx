import styles from "./ScoreBoard.module.css";
import Score from "../Score/Score";

type ScoreBoardProps = {
  scores: {
    playerX: number;
    ties: number;
    playerO: number;
  };
  isPVE?: boolean;
};

const ScoreBoard: React.FC<ScoreBoardProps> = ({ scores, isPVE }) => {
  return (
    <div className={styles.scoreBoard}>
      <Score label="X (P1)" value={scores.playerX} type="playerX" />
      <Score label="TIES" value={scores.ties} type="ties" />
      <Score label={isPVE ? "O (CPU)" : "O (P2)"} value={scores.playerO} type="playerO" />
    </div>
  );
};

export default ScoreBoard;
