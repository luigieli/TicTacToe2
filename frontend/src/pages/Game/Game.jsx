import styles from "./Game.module.css";
import Board from "../../components/Board/Board";
import Title from "../../components/Title/Title";
import LayoutBorder from "../../components/LayoutBorder/LayoutBorder";

export default function Game() {
  return (
    <div className={styles.main}>
      <Board />
      <LayoutBorder />
      <Title />
    </div>
  );
}
