import React from "react";
import styles from "./Square.module.css";

type SquareProps = {
  value: "X" | "O" | null;
  onClick: () => void;
  color: string;
};

const Square: React.FC<SquareProps> = React.memo(
  ({ value, onClick, color }) => {
    return (
      <button className={styles.square} onClick={onClick} style={{ color }}>
        {value}
      </button>
    );
  }
);

export default Square;
