package api 

import "database/sql"

type APIserver struct {
	addr string
	db *sql.DB 
}


func NewAPIServer (addr string , db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

func (s *APIServer) Run() error{
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	return http.ListenAndServe(s.addr, router)
}