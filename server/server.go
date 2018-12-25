package server

import (
	"../api"
	"../conf"
	"net/http"
)

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

func NewServer(config *conf.Config, service *api.PersonService) *Server {
	return &Server{
		Config:        config,
		PersonService: service,
	}
}
