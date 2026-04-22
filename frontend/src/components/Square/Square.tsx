import styles from "./Square.module.css";

type SquareProps = {
  value: string | null;
  onClick: () => void;
};

const Square: React.FC<SquareProps> = ({ value, onClick }) => {
  const valueClass = value === "X" ? styles.x : value === "O" ? styles.o : "";

  return (
    <button className={`${styles.square} ${valueClass}`} onClick={onClick}>
      {value}
    </button>
  );
};

export default Square;
