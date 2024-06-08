package api

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) handleRoot(w http.ResponseWriter, _ *http.Request) {
	res, err := json.MarshalIndent("Hello, in Root Handler", "", "  ")
	if err != nil {
		log.Printf("Error in encoding the response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)

	w.WriteHeader(http.StatusOK)
}
