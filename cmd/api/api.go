package api 

import (
	"database/sql"
	"net/http"
	"github.com/gorilla/mux" // Importing Gorilla Mux for handling routes
)

type APIServer struct {
	addr string
	db   *sql.DB // Database connection
}

// NewAPIServer is a constructor for APIServer. It initializes a new APIServer with the given address and database.
func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr, // Address where the server will run (e.g., ":8080")
		db:   db,   // Database connection
	}
}

// Run starts the HTTP server and listens for incoming requests
func (s *APIServer) Run() error {
	router := mux.NewRouter()                       // Creating a new main router
	subrouter := router.PathPrefix("/api/v1").Subrouter() // Creating a sub-router under /api/v1

	// Creating a user handler and registering user-related routes
	userHandler := user.NewHandler()          // Create a new handler for user routes (like login and register)
	userHandler.RegisterRoutes(subrouter)     // Register user routes with the sub-router

	// Starting the HTTP server at the given address and linking it to the main router
	return http.ListenAndServe(s.addr, router) // Run the server on the given address (e.g., ":8080")
}
