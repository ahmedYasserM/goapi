package main

import (
	"flag"
	"log"

	"github.com/ahmedYasserM/goapi/cmd/api"
	"github.com/ahmedYasserM/goapi/cmd/storage"
)

func main() {
	port := flag.String("port", ":7000", "Listen port")
	flag.Parse()

	db, err := storage.NewPostgres()
	if err != nil {
		log.Fatalf("Error in creating a new postgres instance: %v\n", err)
	}

	server := api.NewServer(*port, db)
	log.Printf("Server is running on port %s\n", *port)

	if err := server.Start(); err != nil {
		log.Fatalf("Error in starting the server: %v\n", err)
	}
}
