package handlers

import (
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) (err error){
	_, err = w.Write([]byte("Hello, in Root Handler"))
	return
}
