import React, { useRef, useState } from "react";
import styles from "./NewGameButton.module.css";

const NewGameButton: React.FC = () => {
  const [isVisible, setIsVisible] = useState(true);

  const onClick = () => setIsVisible(false);

  return isVisible ? (
    <div className={styles.divBlur}>
      <button className={styles.newGameButton} onClick={onClick}>
        Novo Jogo
      </button>
    </div>
  ) : null;
};

export default NewGameButton;
