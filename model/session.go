package model

import (
	"github.com/google/uuid"
)

type Session struct {
	ID     int
	UUID   string
	UserID int
}

func (s *Session) Create() error {
	_, err := db.Exec("INSERT INTO sessions (uuid, user_id) VALUES (?, ?)", uuid.New().String(), s.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) ReadByID() error {
	return db.QueryRow("SELECT uuid, user_id FROM sessions WHERE id = ?", s.ID).
		Scan(&s.UUID, &s.UserID)
}

func (s *Session) ReadByUUID() error {
	return db.QueryRow("SELECT id, user_id FROM sessions WHERE uuid = ?", s.UUID).
		Scan(&s.ID, &s.UserID)
}

func (s *Session) ReadByUserID() error {
	return db.QueryRow("SELECT id, uuid FROM sessions WHERE user_id = ?", s.UserID).
		Scan(&s.ID, &s.UUID)
}

func (s *Session) ReadAll() ([]Session, error) {
	rows, err := db.Query("SELECT id, uuid, user_id FROM sessions")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sessions []Session
	for rows.Next() {
		session := Session{}
		if err := rows.Scan(&session.ID, &session.UUID, &session.UserID); err != nil {
			return nil, err
		}
		sessions = append(sessions, session)
	}
	return sessions, nil
}

func (s *Session) Update() error {
	_, err := db.Exec("UPDATE sessions SET uuid = ?, user_id = ? WHERE id = ?", s.UUID, s.UserID, s.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) Delete() error {
	_, err := db.Exec("DELETE FROM sessions WHERE id = ?", s.ID)
	if err != nil {
		return err
	}
	return nil
}

func (s *Session) DeleteAll() error {
	_, err := db.Exec("DELETE FROM sessions")
	if err != nil {
		return err
	}
	return nil
}
