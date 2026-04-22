import styles from "./Score.module.css";

interface ScoreProps {
  label: string;
  value: number;
  type: 'playerX' | 'playerO' | 'ties';
}

const Score: React.FC<ScoreProps> = ({ label, value, type }) => {
  return (
    <div className={`${styles.scoreItem} ${styles[type]}`}>
      <span className={styles.label}>{label}</span>
      <span className={styles.value}>{value}</span>
    </div>
  );
};

export default Score;
