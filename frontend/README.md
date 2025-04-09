# React + Vite with TypeScript

This template provides a minimal setup to get React working in Vite with TypeScript, HMR, and some ESLint rules.

## Project Structure

- **frontendtypescript/**: Root directory of the project.
- **.gitignore**: Rules for ignoring files and folders that should not be tracked by Git.
- **eslint.config.js**: ESLint configuration for the project, including plugins and specific rules for JavaScript and React.
- **icon.svg**: SVG icon file that can be used as a favicon or elsewhere in the application.
- **index.html**: Main HTML file that loads the React application, containing the basic document structure and an element with the id "root" where the application will mount.
- **package.json**: npm configuration file, including dependencies, scripts, and project information.
- **README.md**: Documentation for the project, including usage instructions and setup information.
- **tsconfig.json**: TypeScript configuration file specifying compiler options and files to include in the compilation.
- **vite.config.ts**: Vite configuration file for the project, including plugins and options specific to the development environment.
- **src/**: Source directory containing the application code.
  - **App.tsx**: Entry point of the React application, importing and rendering the Game component.
  - **main.tsx**: Responsible for initializing the React application, mounting it to the "root" element in index.html.
  - **components/**: Directory containing reusable components.
    - **Board/**: Contains Board.tsx and Board.module.css for the game board.
    - **LayoutBorder/**: Contains LayoutBorder.tsx and LayoutBorder.module.css for decorative borders.
    - **Score/**: Contains Score.tsx and Score.module.css for displaying player scores.
    - **Square/**: Contains Square.tsx and Square.module.css for representing a cell in the game board.
    - **Title/**: Contains Title.tsx and title.module.css for displaying the game title.
  - **pages/**: Directory containing page components.
    - **Game/**: Contains Game.tsx and Game.module.css for the game logic and board components.
  - **styles/**: Directory containing global styles.
    - **App.css**: Global styles for the application.
    - **index.css**: Configuration styles for the body of the application.

## Getting Started

1. Clone the repository.
2. Navigate to the project directory.
3. Install dependencies using `npm install`.
4. Start the development server with `npm run dev`.
5. Open your browser and go to `http://localhost:3000` to see the application in action.

## Expanding the ESLint Configuration

If you are developing a production application, we recommend using TypeScript and enabling type-aware lint rules. Check out the [TypeScript template](https://github.com/vitejs/vite/tree/main/packages/create-vite/template-react-ts) to integrate TypeScript and [`typescript-eslint`](https://typescript-eslint.io) in your project.