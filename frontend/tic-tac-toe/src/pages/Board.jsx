import "../styles/Board.css";

function Board() {
  return (
    <div className="main">
      <div className="main-board"></div>
      <div className="main-text">
        <span className="tic-toe">tic.</span>
        <span className="tac">tac.</span>
        <span className="tic-toe">toe.</span>
      </div>
    </div>
  );
}

export default Board;
