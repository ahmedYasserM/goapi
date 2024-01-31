package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	"github.com/ahmedYasserM/goApi/handlers"
)

type customHandlerFunc func(http.ResponseWriter, *http.Request) error

func (fn customHandlerFunc) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := fn(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	serverPort := flag.String("port", "7000", "Server Port")
	flag.Parse()

	router := mux.NewRouter()

	router.Handle("/", customHandlerFunc(handlers.Root)).Methods(http.MethodGet)
	router.Handle("/posts", customHandlerFunc(handlers.CreatePost)).Methods(http.MethodPost)
	router.Handle("/posts", customHandlerFunc(handlers.ShowAllPosts)).Methods(http.MethodGet)
	router.Handle("/posts/{id}", customHandlerFunc(handlers.ShowPost)).Methods(http.MethodGet)
	router.Handle("/posts/{id}", customHandlerFunc(handlers.ReplacePost)).Methods(http.MethodPut)
	router.Handle("/posts/{id}", customHandlerFunc(handlers.ModifyPost)).Methods(http.MethodPatch)
	router.Handle("/posts/{id}", customHandlerFunc(handlers.DeletePost)).Methods(http.MethodDelete)

	server := &http.Server{
		Addr:    "0.0.0.0:" + *serverPort,
		Handler: router,
	}

	log.Printf("Start listening on port %s\n", *serverPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("%v\n", err)
	}
}
