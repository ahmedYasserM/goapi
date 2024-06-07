package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/ahmedYasserM/goapi/cmd/types"
)

func (s *Server) handleCreatePost(w http.ResponseWriter, r *http.Request) {
	decdoder := json.NewDecoder(r.Body)

	w.Header().Set("Content-Type", "application/json")

	var p types.Post
	err := decdoder.Decode(&p)
	if err != nil {
		log.Printf("Error in decoding the request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.db.CreatePost(&p)
	if err != nil {
		log.Printf("New Post is Added With Id = %d\n", p.Id)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	resutlJson, _ := json.MarshalIndent(struct{ PostId int }{p.Id}, "", "\t")
	fmt.Fprintf(w, "%s\n", resutlJson)
}

func (s *Server) handleShowAllPosts(w http.ResponseWriter, r *http.Request) {
	posts, err := s.db.GetAllPosts()

	if err != nil {
		log.Printf("Error in getting all posts: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	err = encoder.Encode(posts)
	if err != nil {
		log.Printf("Error in encoding the response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleShowPostById(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Error in converting the id to integer: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	p, err := s.db.GetPostById(id)
	if err != nil {
		log.Printf("Error in getting post by id: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonResponse, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		log.Printf("Error in encoding the response: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "%s\n", jsonResponse)

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleModifyPostById(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Error in encoding the response: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	post, err := s.db.GetPostById(id)
	if err != nil {
		log.Printf("Error in getting post by id: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&post)
	if err != nil {
		log.Printf("Error in decoding the request body: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.db.ModifyPostById(id, &post)
	if err != nil {
		log.Printf("Error in modifying the post: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *Server) handleDeletePostById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		log.Printf("Error in getting post by id: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = s.db.DeletePostById(id)
	if err != nil {
		log.Printf("Error in deleting the post: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// func (s *Server) handleReplacePostById(w http.ResponseWriter, r *http.Request) {
// 	w.WriteHeader(http.StatusNotImplemented)
// 	return
// }
