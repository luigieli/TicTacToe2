import styles from "./ScoreBoard.module.css";
import Score from "../Score/Score";

type ScoreBoardProps = {
  scores: {
    playerX: number;
    ties: number;
    playerO: number;
  };
};

const ScoreBoard: React.FC<ScoreBoardProps> = ({ scores }) => {
  return (
    <div className={styles.score}>
      <Score
        player={"JOGADOR X"}
        score={scores.playerX}
        backgroundColor="#48D2FE"
      />
      <Score player={"EMPATE"} score={scores.ties} backgroundColor="#BCDBF9" />
      <Score
        player={"JOGADOR O"}
        score={scores.playerO}
        backgroundColor="#E2BE00"
      />
    </div>
  );
};

export default ScoreBoard;
