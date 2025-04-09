import styles from "./Score.module.css";

const Score: React.FC<{ player: string; score: number }> = ({ player, score }) => {
    return (
        <div className={styles.score}>
            <span className={styles.textScore}>{player}</span>
            <span className={`${styles.textScore} ${styles.textScoreNumber}`}>{score}</span>
        </div>
    );
};

export default Score;