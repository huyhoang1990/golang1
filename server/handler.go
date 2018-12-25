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

func (s *Server) people(w http.ResponseWriter, r *http.Request) {
	people := s.PersonService.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
