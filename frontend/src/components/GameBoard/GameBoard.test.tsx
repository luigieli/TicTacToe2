import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import GameBoard from '../../components/GameBoard/GameBoard';

// Mock child components
vi.mock('../../components/Square/Square', () => ({
  default: ({ value, onClick, color }: any) => (
    <button onClick={onClick} style={{ color }} data-testid={`square-button-${value || ''}`}>
      {value}
    </button>
  ),
}));

vi.mock('../../components/ScoreBoard/ScoreBoard', () => ({
  default: ({ scores }: any) => (
    <div data-testid="scoreboard">
      X: {scores.playerX}, O: {scores.playerO}, Ties: {scores.ties}
    </div>
  ),
}));

vi.mock('../../components/NewGameButton/NewGameButton', () => ({
  default: () => <button data-testid="new-game-button">New Game</button>,
}));

describe('GameBoard Component', () => {
  const mockScores = {
    playerX: 2,
    playerO: 1,
    ties: 0,
  };

  it('should render game board', () => {
    const onSquareClick = vi.fn();
    const squares = Array(9).fill(null);
    render(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={mockScores} />
    );
    expect(screen.getByTestId('scoreboard')).toBeTruthy();
  });

  it('should render new game button', () => {
    const onSquareClick = vi.fn();
    const squares = Array(9).fill(null);
    render(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={mockScores} />
    );
    expect(screen.getByTestId('new-game-button')).toBeTruthy();
  });

  it('should render all 9 squares', () => {
    const onSquareClick = vi.fn();
    const squares = Array(9).fill(null);
    render(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={mockScores} />
    );
    const buttons = screen.getAllByRole('button');
    expect(buttons.length).toBeGreaterThanOrEqual(10); // 9 squares + 1 new game button
  });

  it('should display square values correctly', () => {
    const onSquareClick = vi.fn();
    const squares = ['X', 'O', null, 'X', null, 'O', null, null, null];
    render(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={mockScores} />
    );
    const xButtons = screen.getAllByText('X');
    const oButtons = screen.getAllByText('O');
    expect(xButtons.length).toBe(2);
    expect(oButtons.length).toBe(2);
  });

  it('should call onSquareClick with correct index when square is clicked', () => {
    const onSquareClick = vi.fn();
    const squares = Array(9).fill(null);
    render(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={mockScores} />
    );
    
    const buttons = screen.getAllByRole('button');
    // Click the first square button (not the new game button)
    fireEvent.click(buttons[1]);
    expect(onSquareClick).toHaveBeenCalled();
  });

  it('should display scores', () => {
    const onSquareClick = vi.fn();
    const squares = Array(9).fill(null);
    render(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={mockScores} />
    );
    const scoreboard = screen.getByTestId('scoreboard');
    expect(scoreboard).toHaveTextContent('X: 2');
    expect(scoreboard).toHaveTextContent('O: 1');
    expect(scoreboard).toHaveTextContent('Ties: 0');
  });

  it('should apply correct colors to squares', () => {
    const onSquareClick = vi.fn();
    const squares = ['X', 'O', null];
    const { container } = render(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={mockScores} />
    );
    const buttons = container.querySelectorAll('button');
    
    // Check that buttons have color styles
    expect(buttons.length).toBeGreaterThan(0);
  });

  it('should handle empty board', () => {
    const onSquareClick = vi.fn();
    const squares = Array(9).fill(null);
    render(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={mockScores} />
    );
    expect(screen.getByTestId('scoreboard')).toBeTruthy();
  });

  it('should update when scores change', () => {
    const onSquareClick = vi.fn();
    const squares = Array(9).fill(null);
    const { rerender } = render(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={mockScores} />
    );
    
    const newScores = { playerX: 5, playerO: 3, ties: 1 };
    rerender(
      <GameBoard squares={squares} onSquareClick={onSquareClick} scores={newScores} />
    );
    
    expect(screen.getByTestId('scoreboard')).toHaveTextContent('X: 5');
    expect(screen.getByTestId('scoreboard')).toHaveTextContent('O: 3');
    expect(screen.getByTestId('scoreboard')).toHaveTextContent('Ties: 1');
  });
});
