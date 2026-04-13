import { describe, it, expect } from 'vitest';
import calculateWinner from './CalculateWinner';

describe('calculateWinner', () => {
  it('should return null when board is empty', () => {
    const board = Array(9).fill(null);
    expect(calculateWinner(board)).toBeNull();
  });

  it('should return null when board has no winner', () => {
    const board = ['X', 'O', 'X', 'O', 'X', null, null, null, null];
    expect(calculateWinner(board)).toBeNull();
  });

  it('should detect winner in top row', () => {
    const board = ['X', 'X', 'X', 'O', 'O', null, null, null, null];
    expect(calculateWinner(board)).toBe('X');
  });

  it('should detect winner in middle row', () => {
    const board = ['X', null, null, 'O', 'O', 'O', 'X', null, null];
    expect(calculateWinner(board)).toBe('O');
  });

  it('should detect winner in bottom row', () => {
    const board = ['X', 'O', null, 'O', 'X', null, 'X', 'X', 'X'];
    expect(calculateWinner(board)).toBe('X');
  });

  it('should detect winner in left column', () => {
    const board = ['X', 'O', 'O', 'X', 'O', null, 'X', null, null];
    expect(calculateWinner(board)).toBe('X');
  });

  it('should detect winner in middle column', () => {
    const board = ['X', 'O', 'X', null, 'O', null, null, 'O', 'X'];
    expect(calculateWinner(board)).toBe('O');
  });

  it('should detect winner in right column', () => {
    const board = ['X', 'O', 'X', 'O', 'X', 'O', null, null, 'X'];
    expect(calculateWinner(board)).toBe('X');
  });

  it('should detect winner in left-to-right diagonal', () => {
    const board = ['X', 'O', 'O', null, 'X', 'O', null, null, 'X'];
    expect(calculateWinner(board)).toBe('X');
  });

  it('should detect winner in right-to-left diagonal', () => {
    const board = ['X', 'O', 'O', 'X', 'O', 'X', 'O', null, 'X'];
    expect(calculateWinner(board)).toBe('O');
  });

  it('should return X when X wins', () => {
    const board = ['X', 'X', 'X', 'O', 'O', null, null, null, null];
    expect(calculateWinner(board)).toBe('X');
  });

  it('should return O when O wins', () => {
    const board = ['X', 'O', null, 'X', 'O', null, null, 'O', 'X'];
    expect(calculateWinner(board)).toBe('O');
  });

  it('should return null for almost winning scenario', () => {
    const board = ['X', 'X', null, 'O', 'O', null, null, null, null];
    expect(calculateWinner(board)).toBeNull();
  });
});
