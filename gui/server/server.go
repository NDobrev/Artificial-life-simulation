package server

import (
	"net/http"
)

type Server struct {
}

func (s *Server) Run() {
	http.ListenAndServe(":8090", nil)
}

func (s *Server) RegisterFunc(url string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(url, handler)
}
