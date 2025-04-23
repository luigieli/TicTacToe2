<<<<<<< HEAD
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
=======
import React, { useRef } from "react";
import styles from "./NewGameButton.module.css";

const NewGameButton: React.FC = () => {
  const buttonRef = useRef<HTMLButtonElement>(null);
  
  const onClick = () => {
    if(buttonRef.current && buttonRef.current.parentElement) {
      const divParent = buttonRef.current.parentElement as HTMLDivElement;
      divParent.style.display="none"
    }
  }
  
  return (
    <div className={styles.divBlur}>
      <button ref={buttonRef} className={styles.newGameButton} onClick={onClick}>Novo Jogo</button>
    </div>
  );
>>>>>>> 699ddaa40560219fff5386a7d48bccd954f454bb
};

export default NewGameButton;
