package storage

import (
	"os"

	"github.com/ahmedYasserM/goapi/cmd/types"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres() (*Postgres, error) {
	db, err := sqlx.Open("postgres", os.Getenv("DB_URL"))
	return &Postgres{db: db}, err
}

func (psg *Postgres) CreatePost(post *types.Post) (err error) {

	stmt, err := psg.db.Prepare("INSERT INTO posts (content, author) VALUES ($1, $2) RETURNING id;")

	if err != nil {
		return err
	}

	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)

	return err
}

func (psg *Postgres) GetAllPosts() (posts []types.Post, err error) {
	rows, err := psg.db.Queryx("SELECT * FROM posts ORDER BY id DESC;")

	if err != nil {
		return nil, err
	}

	var post types.Post
	for rows.Next() {
		if err = rows.StructScan(&post); err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (psg *Postgres) GetPostById(id int) (post types.Post, err error) {
	err = psg.db.QueryRowx("SELECT * FROM posts WHERE id = $1;", id).StructScan(&post)

	return post, err
}

func (psg *Postgres) ModifyPostById(id int, post *types.Post) (err error) {
	_, err = psg.db.Exec("UPDATE posts SET author = $2, content = $3 WHERE id = $1;", post.Id, post.Author, post.Content)

	return err
}

func (psg *Postgres) DeletePostById(id int) (err error) {
	_, err = psg.db.Exec("DELETE FROM posts WHERE id = $1;", id)

	return err
}
