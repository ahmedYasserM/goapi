package api

import (
	"github.com/ahmedYasserM/goapi/storage"
	"net/http"
)

type Server struct {
	port string
	db   storage.Storage
}

func NewServer(port string, db storage.Storage) *Server {
	return &Server{
		port: port,
		db:   db,
	}

}

func (s *Server) Start() error {
	http.HandleFunc("GET /", s.handleRoot)
	http.HandleFunc("POST /posts", s.handleCreatePost)
	http.HandleFunc("GET /posts", s.handleShowAllPosts)
	http.HandleFunc("GET /posts/{id}", s.handleShowPostById)
	http.HandleFunc("PATCH /posts/{id}", s.handleModifyPostById)
	http.HandleFunc("DELETE /posts/{id}", s.handleDeletePostById)

	return http.ListenAndServe(s.port, nil)
}
