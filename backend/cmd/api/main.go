package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"tictactoe/internal/api/handlers"
	"tictactoe/internal/repository"
	"tictactoe/internal/service"
)

func main() {
	// Dependency Injection
	aiSvc := service.NewAIService()

	tictactoe2Repo := repository.NewGameRepository()
	tictactoe2Svc := service.NewGameService(tictactoe2Repo, aiSvc)
	tictactoe2Handler := handlers.NewTicTacToe2Handler(tictactoe2Svc)

	standardRepo := repository.NewStandardGameRepository()
	standardSvc := service.NewStandardGameService(standardRepo, aiSvc)
	standardHandler := handlers.NewStandardGameHandler(standardSvc)

	// Router setup
	r := chi.NewRouter()

	// Standard Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	// Routes
	r.Route("/games", func(r chi.Router) {
		r.Route("/tictactoe2", func(r chi.Router) {
			r.Post("/", tictactoe2Handler.CreateGame)
			r.Get("/{id}", tictactoe2Handler.GetGameState)
			r.Post("/{id}/move", tictactoe2Handler.MakeMove)
		})

		r.Route("/standard", func(r chi.Router) {
			r.Post("/", standardHandler.CreateGame)
			r.Get("/{id}", standardHandler.GetGameState)
			r.Post("/{id}/move", standardHandler.MakeMove)
		})
	})

	port := ":8080"
	fmt.Printf("Tic-Tac-Toe-2 Backend starting on %s...\n", port)
	if err := http.ListenAndServe(port, r); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
