package storage

import "github.com/ahmedYasserM/goapi/cmd/types"

type Storage interface {
	CreatePost(post *types.Post) error
	GetAllPosts() ([]types.Post, error)
	GetPostById(id int) (types.Post, error)
	ModifyPostById(id int, post *types.Post) error
	DeletePostById(id int) error
}
