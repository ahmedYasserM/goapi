package db

import (
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var db *sqlx.DB

func init() {
	var err error
	// db, err = sqlx.Open("postgres", "postgres://postgres:123@localhost:5432/goweb?sslmode=disable")
	db, err = sqlx.Open("postgres", os.Getenv("DB_URL"))

	if err != nil {
		log.Fatalf("%v\n", err)
	}
}

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   string    `json:"author"`
}


// TODO: Create
func (post *Post) Create() (err error) {
	stmt, err := db.Prepare("INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id;")
	if err != nil {
		return
	}

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)

	return
}

// TODO: Read(one)
func GetPost(id int) (p Post, err error) {
	err = db.QueryRowx("SELECT * FROM posts WHERE id = $1;", id).StructScan(&p)

	return
}

// TODO: Read(all)
func GetPosts() (posts []Post, err error) {
	rows, err := db.Queryx("SELECT * FROM posts ORDER BY id DESC;")

	if err != nil {
		return
	}

	var p Post
	for rows.Next() {
		if err = rows.StructScan(&p); err != nil {
			return
		}

		posts = append(posts, p)
	}

	return
}

// TODO: Update
func (p *Post) Modify() (err error) {
	_, err = db.Exec("UPDATE posts SET author = $2, content = $3 WHERE id = $1;", p.Id, p.Author, p.Content)

	return
}

// TODO: Delete
func (p *Post) Delete() (err error) {
	_, err = db.Exec("DELETE FROM posts WHERE id = $1;", p.Id)

	return
}
