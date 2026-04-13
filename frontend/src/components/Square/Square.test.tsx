import { describe, it, expect, vi } from 'vitest';
import { render, screen, fireEvent } from '@testing-library/react';
import Square from '../../components/Square/Square';

describe('Square Component', () => {
  it('should render a button', () => {
    const onClick = vi.fn();
    render(<Square value={null} onClick={onClick} color="" />);
    expect(screen.getByRole('button')).toBeTruthy();
  });

  it('should display the value', () => {
    const onClick = vi.fn();
    render(<Square value="X" onClick={onClick} color="" />);
    expect(screen.getByText('X')).toBeTruthy();
  });

  it('should call onClick when clicked', () => {
    const onClick = vi.fn();
    render(<Square value="X" onClick={onClick} color="" />);
    fireEvent.click(screen.getByRole('button'));
    expect(onClick).toHaveBeenCalledOnce();
  });

  it('should apply color style', () => {
    const onClick = vi.fn();
    const { container } = render(
      <Square value="X" onClick={onClick} color="#72CFF9" />
    );
    const button = container.querySelector('button');
    expect(button?.style.color).toBe('rgb(114, 207, 249)');
  });

  it('should display null as empty string', () => {
    const onClick = vi.fn();
    const { container } = render(<Square value={null} onClick={onClick} color="" />);
    const button = container.querySelector('button');
    expect(button?.textContent).toBe('');
  });

  it('should display O with yellow color', () => {
    const onClick = vi.fn();
    const { container } = render(
      <Square value="O" onClick={onClick} color="#E2BE00" />
    );
    const button = container.querySelector('button');
    expect(button?.style.color).toBe('rgb(226, 190, 0)');
    expect(screen.getByText('O')).toBeTruthy();
  });

  it('should handle multiple clicks', () => {
    const onClick = vi.fn();
    render(
      <Square value="X" onClick={onClick} color="#72CFF9" />
    );
    fireEvent.click(screen.getByRole('button'));
    fireEvent.click(screen.getByRole('button'));
    expect(onClick).toHaveBeenCalledTimes(2);
  });
});
