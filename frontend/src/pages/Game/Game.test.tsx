import { describe, it, expect, vi, beforeEach } from 'vitest';
import { render, screen, fireEvent, waitFor } from '@testing-library/react';
import Game from '../../pages/Game/Game';
import * as api from '../../utils/api';

// Mock the API
vi.mock('../../utils/api', () => ({
  createStandardGame: vi.fn(),
  makeStandardMove: vi.fn(),
}));

// Mock the components used in Game
vi.mock('../../components/GameBoard/GameBoard', () => ({
  default: ({ squares, onSquareClick, scores }: any) => (
    <div data-testid="game-board">
      <div data-testid="scores">
        X: {scores.playerX}, O: {scores.playerO}, Ties: {scores.ties}
      </div>
      <div data-testid="squares">
        {squares.map((square: any, index: number) => (
          <button
            key={index}
            data-testid={`square-${index}`}
            onClick={() => onSquareClick(index)}
          >
            {square || ''}
          </button>
        ))}
      </div>
    </div>
  ),
}));

vi.mock('../../components/Title/Title', () => ({
  default: () => <div data-testid="title">Title</div>,
}));

vi.mock('../../components/FrameBorder/FrameBorder', () => ({
  default: () => <div data-testid="frame-border">FrameBorder</div>,
}));

describe('Game Component', () => {
  beforeEach(() => {
    vi.clearAllMocks();
    (api.createStandardGame as any).mockResolvedValue('test-game-id');
  });

  it('should render game board with initial state', async () => {
    render(<Game />);
    await waitFor(() => {
      expect(screen.getByTestId('game-board')).toBeTruthy();
    });
  });

  it('should render title and frame border', async () => {
    render(<Game />);
    await waitFor(() => {
      expect(screen.getByTestId('title')).toBeTruthy();
      expect(screen.getByTestId('frame-border')).toBeTruthy();
    });
  });

  it('should have initial scores at 0', async () => {
    render(<Game />);
    await waitFor(() => {
      expect(screen.getByTestId('scores')).toHaveTextContent('X: 0, O: 0, Ties: 0');
    });
  });

  it('should update board when square is clicked', async () => {
    (api.makeStandardMove as any).mockResolvedValue({
      id: 'test-game-id',
      mode: 'PVP',
      board: ['X', '', '', '', '', '', '', '', ''],
      current_player: 'O',
      winner: '',
      is_game_over: false,
    });

    render(<Game />);
    await waitFor(() => screen.getByTestId('square-0'));
    
    const square0 = screen.getByTestId('square-0');
    fireEvent.click(square0);
    
    await waitFor(() => {
      expect(square0.textContent).toBe('X');
    });
  });

  it('should alternate players on each move', async () => {
    (api.makeStandardMove as any)
      .mockResolvedValueOnce({
        id: 'test-game-id',
        mode: 'PVP',
        board: ['X', '', '', '', '', '', '', '', ''],
        current_player: 'O',
        winner: '',
        is_game_over: false,
      })
      .mockResolvedValueOnce({
        id: 'test-game-id',
        mode: 'PVP',
        board: ['X', 'O', '', '', '', '', '', '', ''],
        current_player: 'X',
        winner: '',
        is_game_over: false,
      })
      .mockResolvedValueOnce({
        id: 'test-game-id',
        mode: 'PVP',
        board: ['X', 'O', 'X', '', '', '', '', '', ''],
        current_player: 'O',
        winner: '',
        is_game_over: false,
      });

    render(<Game />);
    await waitFor(() => screen.getByTestId('square-0'));

    fireEvent.click(screen.getByTestId('square-0'));
    await waitFor(() => expect(screen.getByTestId('square-0').textContent).toBe('X'));

    fireEvent.click(screen.getByTestId('square-1'));
    await waitFor(() => expect(screen.getByTestId('square-1').textContent).toBe('O'));

    fireEvent.click(screen.getByTestId('square-2'));
    await waitFor(() => expect(screen.getByTestId('square-2').textContent).toBe('X'));
  });

  it('should not allow clicking on occupied square', async () => {
    (api.makeStandardMove as any).mockResolvedValue({
      id: 'test-game-id',
      mode: 'PVP',
      board: ['X', '', '', '', '', '', '', '', ''],
      current_player: 'O',
      winner: '',
      is_game_over: false,
    });

    render(<Game />);
    await waitFor(() => screen.getByTestId('square-0'));

    const square0 = screen.getByTestId('square-0');
    fireEvent.click(square0);
    
    await waitFor(() => expect(square0.textContent).toBe('X'));

    // Second click - should not call makeMove again
    fireEvent.click(square0);
    expect(api.makeStandardMove).toHaveBeenCalledTimes(1);
  });

  it('should render 9 squares', async () => {
    render(<Game />);
    await waitFor(() => {
      for (let i = 0; i < 9; i++) {
        expect(screen.getByTestId(`square-${i}`)).toBeTruthy();
      }
    });
  });
});
