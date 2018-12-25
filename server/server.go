package server

import (
	"../api"
	"../conf"
	"encoding/json"
	"net/http"
)

type Server struct {
	Config        *conf.Config
	PersonService *api.PersonService
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/people", s.people)

	return mux
}

func (s *Server) Run() {
	httpServer := &http.Server{
		Addr:    ":" + s.Config.Port,
		Handler: s.Handler(),
	}

	httpServer.ListenAndServe()
}

func (s *Server) people(w http.ResponseWriter, r *http.Request) {
	people := s.PersonService.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}

func NewServer(config *conf.Config, service *api.PersonService) *Server {
	return &Server{
		Config:        config,
		PersonService: service,
	}
}
