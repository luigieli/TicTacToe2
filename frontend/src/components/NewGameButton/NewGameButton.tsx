import React, { useState } from "react";
import styles from "./NewGameButton.module.css";
import { STRINGS } from "../../constants";

const NewGameButton: React.FC = () => {
  const [isVisible, setIsVisible] = useState(true);

  const onClick = () => setIsVisible(false);

  return isVisible ? (
    <div className={styles.divBlur}>
      <button className={styles.newGameButton} onClick={onClick}>
        {STRINGS.NEW_GAME}
      </button>
    </div>
  ) : null;
};

export default NewGameButton;
