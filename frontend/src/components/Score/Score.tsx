import styles from "./Score.module.css";

const Score: React.FC<{ player: string; score: number, backgroundColor: string}> = ({
  player,
  score,
  backgroundColor
}) => {
  return (
    <div className={styles.score} style={{backgroundColor}}>
      <span className={styles.textScore}>{player}</span>
      <span className={`${styles.textScore} ${styles.textScoreNumber}`}>
        {score}
      </span>
    </div>
  );
};

export default Score;
