import { useState } from "react";
import "./styles/App.css";
import Game from "./pages/Game/Game";
import Menu from "./pages/Menu/Menu";
import { GameMode } from "./utils/api";

export default function App() {
  const [gameMode, setGameMode] = useState<"MENU" | GameMode>("MENU");

  const handleSelectMode = (mode: GameMode) => {
    setGameMode(mode);
  };

  const handleBackToMenu = () => {
    setGameMode("MENU");
  };

  return (
    <>
      {gameMode === "MENU" ? (
        <Menu onSelectMode={handleSelectMode} />
      ) : (
        <Game mode={gameMode} onBack={handleBackToMenu} />
      )}
    </>
  );
}
