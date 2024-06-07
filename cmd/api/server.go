package api

import (
	"github.com/ahmedYasserM/goapi/cmd/storage"
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
	router := http.NewServeMux()

	router.HandleFunc("GET /", s.handleRoot)
	router.HandleFunc("POST /posts", s.handleCreatePost)
	router.HandleFunc("GET /posts", s.handleShowAllPosts)
	router.HandleFunc("GET /posts/{id}", s.handleShowPostById)
	router.HandleFunc("PATCH /posts/{id}", s.handleModifyPostById)
	router.HandleFunc("DELETE /posts/{id}", s.handleDeletePostById)

	server := http.Server{Addr: s.port, Handler: router}

	return server.ListenAndServe()
}
