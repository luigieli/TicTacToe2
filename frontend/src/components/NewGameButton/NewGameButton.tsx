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
};

export default NewGameButton;
