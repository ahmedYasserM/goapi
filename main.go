package main

import (
	"flag"
	"log"

	"github.com/ahmedYasserM/goapi/api"
	"github.com/ahmedYasserM/goapi/storage"
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
	log.Fatalf("%s\n", server.Start())
}
