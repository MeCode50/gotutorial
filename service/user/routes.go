package user

import (
	"net/http"
	"github.com/gorilla/mux" // Importing Gorilla Mux for routing
)

// Handler is a struct that will hold methods to handle user-related routes like login and registration.
type Handler struct {}

// NewHandler is a constructor that returns a pointer to a new Handler. 
// It is used to create a handler object that can register routes and handle requests.
func NewHandler() *Handler {
	return &Handler{} // Return an instance of Handler (empty for now)
}

// RegisterRoutes registers user-related routes like login and registration.
// It uses Gorilla Mux's router to handle the incoming HTTP requests.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")    // Handle POST requests to /login with the handleLogin function
	router.HandleFunc("/register", h.handleRegister).Methods("POST") // Handle POST requests to /register with the handleRegister function
}

// handleLogin handles the login functionality when a POST request is made to /login.
// For now, this is empty, but in the future, it will handle user authentication.
func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	// This function will handle user login logic in the future.
}

// handleRegister handles the registration functionality when a POST request is made to /register.
// For now, this is empty, but in the future, it will handle user registration.
func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	// This function will handle user registration logic in the future.
}
