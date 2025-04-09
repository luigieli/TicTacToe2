import styles from "./Score.module.css"

export default function Score(){
    return <div className={styles.score}>
        <span className={styles.textScore}>PLAYER X</span>
        <span className={`${styles.textScore} ${styles.textScoreNumber}`}>1</span>
    </div>
}