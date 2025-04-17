function calculateWinner(board: (string | null)[]): string | null {
    const lines = [
      [0, 1, 2],
      [3, 4, 5],
      [6, 7, 8], // linhas
      [0, 3, 6],
      [1, 4, 7],
      [2, 5, 8], // colunas
      [0, 4, 8],
      [2, 4, 6], // diagonais
    ];
  
    for (const [a, b, c] of lines) {
      if (board[a] && board[a] === board[b] && board[a] === board[c]) {
        return board[a];
      }
    }
  
    return null;
  }

export default calculateWinner;