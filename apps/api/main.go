package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	// auth
	r.Post("/api/auth/signup", handleSignup)
	r.Post("/api/auth/login", handleLogin)

	// protected example
	r.Group(func(pr chi.Router) {
		pr.Use(JWTMiddleware())
		pr.Get("/api/me", handleMe)
	})

	// in main.go after auth routes
	r.Group(func(pr chi.Router) {
		pr.Use(JWTMiddleware())
		pr.Post("/api/documents", handleCreateDocument)
	})

	r.Group(func(pr chi.Router) {
		pr.Use(JWTMiddleware())
		pr.Post("/api/review/{document_id}/start", handleStartReview)
		pr.Get("/api/review/{document_id}/findings", handleGetFindings)
	})

	r.Group(func(pr chi.Router) {
		pr.Use(JWTMiddleware())
		pr.Post("/api/matters", handleCreateMatter)
		pr.Get("/api/matters", handleListMatters)
	})

	r.Group(func(pr chi.Router) {
		pr.Use(JWTMiddleware())
		pr.Get("/api/search", handleSearch)
	})

	log.Println("api listening on :8080")
	http.ListenAndServe(":8080", r)
}
