
# RFC 001: Tic-Tac-Toe 2 - Frontend Architecture

**Status:** Draft
**Date:** 2026-03-16

## Introduction

This RFC outlines the initial frontend architecture for the "Tic-Tac-Toe 2" project. The goal is to establish a common understanding of the project's structure, development plan, and technical stack, ensuring that all team members are aligned.

## Project Goals

"Tic-Tac-Toe 2" is an advanced version of the classic Tic-Tac-Toe game. It features a 3x3 grid of smaller Tic-Tac-Toe boards, creating a more complex and strategic gameplay experience.

### Gameplay Rules

1.  **Players:** The game is played by two players, "X" and "O".
2.  **Turns:** Players take turns placing their marks on the board.
3.  **Winning a Small Board:** A player wins a small board by getting three of their marks in a row (horizontally, vertically, or diagonally).
4.  **Winning the Game:** A player wins the overall game by winning three small boards in a row (horizontally, vertically, or diagonally).
5.  **Gameplay Flow:**
    *   The first player can choose any square on any of the nine boards.
    *   The position of the mark on the small board determines the next board the opponent can play on. For example, if a player places their mark in the top-right square of a small board, the next player must play on the top-right board.
    *   If a player is sent to a board that has already been won or is full, they can choose any other board to play on.

### Features

*   **Single Player vs. AI:** A player can play against an AI opponent that uses the Minimax algorithm.
*   **Multiplayer:** Two players can play against each other locally.
*   **Tournament Mode:** A future feature that will allow multiple players to compete in a tournament.

## Development Plan

The development of the frontend will be divided into three main stages:

### Stage 1: Button and Basic Components

This stage focuses on creating the basic building blocks of the UI, such as buttons, and other simple components.

### Stage 2: Single Tic-Tac-Toe Board

This stage involves building a standard 3x3 Tic-Tac-Toe board with the complete game logic for a single game. This includes:

*   Placing marks on the board.
*   Detecting a winner or a draw.
*   Resetting the game.

*This is the current stage of the project. The basic components and a single board are partially implemented.*

### Stage 3: The "Big Board" and Advanced Rules

This is the final stage, where the "Tic-Tac-Toe 2" game is fully implemented. This includes:

*   Creating a "Big Board" component that consists of a 3x3 grid of the single Tic-Tac-Toe boards.
*   Implementing the advanced rules that link the boards together.
*   Integrating the AI opponent.
*   Implementing the multiplayer and tournament modes.

## Current Tech Stack

The following technologies are currently being used in the project:

*   **Framework:** React 19
*   **Language:** TypeScript
*   **Build Tool:** Vite
*   **Styling:** CSS Modules
*   **Linting:** ESLint

This RFC is a living document and will be updated as the project evolves.
