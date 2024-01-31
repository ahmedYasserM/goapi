package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/ahmedYasserM/goApi/db"
)

type result struct {
	PostId int `json:"postId"`
}

// TODO: Create
func CreatePost(w http.ResponseWriter, r *http.Request) (err error) {
	decdoder := json.NewDecoder(r.Body)

	defer func() {
		if err = r.Body.Close(); err != nil {
			return
		}
	}()

	w.Header().Set("Content-Type", "application/json")

	var p db.Post
	err = decdoder.Decode(&p)
	if err != nil && err != io.EOF {
		return
	}

	err = p.Create()
	log.Printf("New Post is Added With Id = %d\n", p.Id)

	w.WriteHeader(http.StatusOK)

	resutlJson, _ := json.MarshalIndent(result{PostId: p.Id}, "", "\t")
	fmt.Fprintf(w, "%s\n", resutlJson)

	return
}

// TODO: Read(all)
func ShowAllPosts(w http.ResponseWriter, r *http.Request) (err error) {
	posts, err := db.GetPosts()
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(posts)
	if err != nil {
		return
	}

	w.WriteHeader(http.StatusOK)
	return
}

// TODO: Read(one) (err error)
func ShowPost(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "application/json")

	p, err := db.GetPost(id)
	if err != nil {
		return
	}

	jsonResponse, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		return
	}

	fmt.Fprintf(w, "%s\n", jsonResponse)

	w.WriteHeader(http.StatusOK)
	return
}

// TODO: Update(Replace)
func ReplacePost(w http.ResponseWriter, r *http.Request) (err error) {
	w.WriteHeader(http.StatusNotImplemented)
	return
}

// TODO: Update(Modify)
func ModifyPost(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}

	p, err := db.GetPost(id)
	if err != nil {
		return
	}

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&p)
	if err != nil {
		return
	}

	err = p.Modify()
	w.WriteHeader(http.StatusOK)

	return
}

// TODO: Delete
func DeletePost(w http.ResponseWriter, r *http.Request) (err error) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		return
	}

	p, err := db.GetPost(id)
	if err != nil {
		return
	}

	err = p.Delete()
	w.WriteHeader(http.StatusOK)

	return
}

