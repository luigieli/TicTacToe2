import styles from "./ScoreBoard.module.css";
import Score from "../Score/Score";
import { STRINGS, COLORS } from "../../constants";

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
        player={STRINGS.PLAYER_X}
        score={scores.playerX}
        backgroundColor={COLORS.PLAYER_X}
      />
      <Score player={STRINGS.TIE} score={scores.ties} backgroundColor="#BCDBF9" />
      <Score
        player={STRINGS.PLAYER_O}
        score={scores.playerO}
        backgroundColor={COLORS.PLAYER_O}
      />
    </div>
  );
};

export default ScoreBoard;
