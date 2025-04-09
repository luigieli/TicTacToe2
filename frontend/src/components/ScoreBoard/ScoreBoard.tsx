import styles from "./ScoreBoard.module.css"
import Score from "../Score/Score";

export default function ScoreBoard() {
  return (
    <div className={styles.score}>
      <Score player={"PLAYER X"} score={0} />
      <Score player={"DRAW"} score={0} />
      <Score player={"PLAYER O"} score={0} />
    </div>
  );
}
