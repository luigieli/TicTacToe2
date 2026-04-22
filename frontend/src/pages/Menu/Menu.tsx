import React, { useState } from 'react';
import styles from './Menu.module.css';
import { GameMode } from '../../utils/api';

interface MenuProps {
  onSelectMode: (mode: GameMode) => void;
}

const Menu: React.FC<MenuProps> = ({ onSelectMode }) => {
  const [showDifficulty, setShowDifficulty] = useState(false);

  if (showDifficulty) {
    return (
      <div className={styles.container}>
        <main className={`${styles.menuCard} ${styles.slideIn}`}>
          <header className={styles.title}>
            <span className={styles.subtitle}>Escolha o Nível</span>
            <h1 className={styles.titleMain}>
              <span className={styles.tic}>difi.</span>
              <span className={styles.tac}>cul.</span>
              <span className={styles.toe}>dade</span>
            </h1>
          </header>

          <nav className={styles.buttonGroup}>
            <button 
              className={`${styles.menuButton} ${styles.pveEasy}`}
              onClick={() => onSelectMode('PVE_EASY')}
            >
              Iniciante
              <span className={styles.buttonIcon}>🐣</span>
            </button>
            
            <button 
              className={`${styles.menuButton} ${styles.pveMedium}`}
              onClick={() => onSelectMode('PVE_MEDIUM')}
            >
              Intermediário
              <span className={styles.buttonIcon}>🧠</span>
            </button>

            <button 
              className={`${styles.menuButton} ${styles.pveHard}`}
              onClick={() => onSelectMode('PVE_HARD')}
            >
              Mestre (IA)
              <span className={styles.buttonIcon}>⚡</span>
            </button>

            <button 
              className={styles.backButton}
              onClick={() => setShowDifficulty(false)}
            >
              ← Voltar ao Início
            </button>
          </nav>
        </main>
      </div>
    );
  }

  return (
    <div className={styles.container}>
      <main className={styles.menuCard}>
        <header className={styles.title}>
          <span className={styles.subtitle}>Welcome to</span>
          <h1 className={styles.titleMain}>
            <span className={styles.tic}>tic.</span>
            <span className={styles.tac}>tac.</span>
            <span className={styles.toe}>toe.</span>
          </h1>
        </header>

        <nav className={styles.buttonGroup}>
          <button 
            className={`${styles.menuButton} ${styles.pvp}`}
            onClick={() => onSelectMode('PVP')}
          >
            PVP Local
            <span className={styles.buttonIcon}>👥</span>
          </button>
          
          <button 
            className={`${styles.menuButton} ${styles.pve}`}
            onClick={() => setShowDifficulty(true)}
          >
            Versus CPU
            <span className={styles.buttonIcon}>🤖</span>
          </button>

          <div className={styles.futureOptions}>
            <button className={`${styles.menuButton} ${styles.disabled}`} disabled>
              Online Battle
              <span>Coming Soon</span>
            </button>
            <button className={`${styles.menuButton} ${styles.disabled}`} disabled>
              Ultimate 9x9
              <span>Locked</span>
            </button>
          </div>
        </nav>
      </main>
    </div>
  );
};

export default Menu;
