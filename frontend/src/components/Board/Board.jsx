import styles from "./Board.module.css"
import Square from "../Square/Square";
import Score from "../Score/Score";
export default function Board() {
  return (
    <div className={styles.mainGame}>
      <div className={styles.score}>
        <Score/>
        <Score/>
        <Score/>
      </div>
      <div className={styles.board}>
        <Square/>
      </div>
    </div>
  );
}
