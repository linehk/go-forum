package model

import (
	"github.com/google/uuid"
)

type Post struct {
	ID       int
	UUID     string
	Content  string
	UserID   int
	ThreadID int
}

func (p *Post) Create() error {
	_, err := db.Exec("INSERT INTO posts (uuid, content, user_id, thread_id) VALUES (?, ?, ?, ?)",
		uuid.New().String(), p.Content, p.UserID, p.ThreadID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) ReadByID() error {
	return db.QueryRow("SELECT uuid, content, user_id, thread_id FROM posts WHERE id = ?", p.ID).
		Scan(&p.UUID, &p.Content, &p.UserID, &p.ThreadID)
}

func (p *Post) ReadByUUID() error {
	return db.QueryRow("SELECT id, content, user_id, thread_id FROM posts WHERE uuid = ?", p.UUID).
		Scan(&p.ID, &p.Content, &p.UserID, &p.ThreadID)
}

func (p *Post) ReadByUserID() ([]Post, error) {
	rows, err := db.Query("SELECT id, uuid, content, user_id, thread_id FROM posts WHERE user_id = ?", p.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.ID, &post.UUID, &post.Content, &post.UserID, &post.ThreadID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *Post) ReadByThreadID() ([]Post, error) {
	rows, err := db.Query("SELECT id, uuid, content, user_id, thread_id FROM posts WHERE thread_id = ?", p.ThreadID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.ID, &post.UUID, &post.Content, &post.UserID, &post.ThreadID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *Post) CountByThreadID() (int, error) {
	var c int
	if err := db.QueryRow("SELECT COUNT(*) FROM posts WHERE thread_id = ?", p.ThreadID).Scan(&c); err != nil {
		return 0, err
	}
	return c, nil
}

func (p *Post) ReadAll() ([]Post, error) {
	rows, err := db.Query("SELECT id, uuid, content, user_id, thread_id FROM posts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var posts []Post
	for rows.Next() {
		post := Post{}
		if err := rows.Scan(&post.ID, &post.UUID, &post.Content, &post.UserID, &post.ThreadID); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (p *Post) Update() error {
	_, err := db.Exec("UPDATE posts SET uuid = ?, content = ?, user_id = ?, thread_id = ?",
		p.UUID, p.Content, p.UserID, p.ThreadID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) Delete() error {
	_, err := db.Exec("DELETE FROM posts WHERE id = ?", p.ID)
	if err != nil {
		return err
	}
	return nil
}

func (p *Post) DeleteAll() error {
	_, err := db.Exec("DELETE FROM posts")
	if err != nil {
		return err
	}
	return nil
}
