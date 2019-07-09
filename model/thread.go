package model

import (
	"time"

	"github.com/google/uuid"
)

type Thread struct {
	ID        int
	CreatedAt time.Time
	UUID      string
	Subject   string
	UserID    int
}

func (t *Thread) Create() error {
	_, err := db.Exec("INSERT INTO threads (created_at, uuid, subject, user_id) VALUES (?, ?, ?, ?)",
		time.Now(), uuid.New().String(), t.Subject, t.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (t *Thread) ReadByID() error {
	return db.QueryRow("SELECT created_at, uuid, subject, user_id FROM threads WHERE id = ?").
		Scan(&t.CreatedAt, &t.UUID, &t.Subject, &t.UserID)
}

func (t *Thread) ReadByUUID() error {
	return db.QueryRow("SELECT id, created_at, subject, user_id FROM threads WHERE uuid = ?").
		Scan(&t.ID, &t.CreatedAt, &t.Subject, &t.UserID)
}

func (t *Thread) ReadBySubject() error {
	return db.QueryRow("SELECT id, created_at, uuid, user_id FROM threads WHERE user_id = ?").
		Scan(&t.ID, &t.CreatedAt, &t.UUID, &t.UserID)
}

func (t *Thread) ReadByUserID() ([]Thread, error) {
	rows, err := db.Query("SELECT id, created_at, uuid, subject, user_id FROM threads WHERE user_id = ?", t.UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var threads []Thread
	for rows.Next() {
		thread := Thread{}
		if err := rows.Scan(&thread.ID, &thread.CreatedAt, &thread.UUID, &thread.Subject, &thread.UserID); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}
	return threads, nil
}

func (t *Thread) ReadAll() ([]Thread, error) {
	rows, err := db.Query("SELECT id, created_at, uuid, subject, user_id FROM threads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var threads []Thread
	for rows.Next() {
		thread := Thread{}
		if err := rows.Scan(&thread.ID, &thread.CreatedAt, &thread.UUID, &thread.Subject, &thread.UserID); err != nil {
			return nil, err
		}
		threads = append(threads, thread)
	}
	return threads, nil
}

func (t *Thread) Update() error {
	_, err := db.Exec("UPDATE threads SET uuid = ?, subject = ?, user_id = ? WHERE id = ?",
		t.UUID, t.Subject, t.UserID, t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *Thread) Delete() error {
	_, err := db.Exec("DELETE FROM threads WHERE id = ?", t.ID)
	if err != nil {
		return err
	}
	return nil
}

func (t *Thread) DeleteAll() error {
	_, err := db.Exec("DELETE FROM threads")
	if err != nil {
		return err
	}
	return nil
}
