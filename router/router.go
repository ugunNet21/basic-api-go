// package router

// import (
// 	"github.com/api/handler"
// 	"github.com/api/middleware"
// 	"github.com/api/repository"
// 	"github.com/go-chi/chi"
// )

// // Your router configurations go here

// func NewRouter(db *repository.Database) *chi.Mux {
// 	r := chi.NewRouter()

// 	// Middleware
// 	r.Use(middleware.LoggingMiddleware)

// 	// Routes
// 	r.Get("/books", handler.GetAllBooks)
// 	// Add more routes as needed

// 	return r
// }
