package user 


type Handler struct {}

func Handler() *Handler {
	return &Handler{}
}

func (h *handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin). Methods("POST")
	router.HandleFunc("/register", h.handleLogin). Methods("POST")

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

