package model

type Hint struct {
	ID       int
	Question string
	Answer   string
	UserID   int
}

func (h *Hint) Create() error {
	_, err := db.Exec("INSERT INTO hints (question, answer, user_id) VALUES (?, ?, ?)",
		h.Question, h.Answer, h.UserID)
	if err != nil {
		return err
	}
	return nil
}

func (h *Hint) ReadByID() error {
	return db.QueryRow("SELECT question, answer, user_id FROM hints WHERE id = ?", h.ID).
		Scan(&h.Question, &h.Answer, &h.UserID)
}

func (h *Hint) ReadByQuestion() ([]Hint, error) {
	rows, err := db.Query("SELECT id, question, answer, user_id FROM hints WHERE question = ?", h.Question)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var hints []Hint
	for rows.Next() {
		hint := Hint{}
		if err := rows.Scan(&hint.ID, &hint.Question, &hint.Answer, &hint.UserID); err != nil {
			return nil, err
		}
		hints = append(hints, hint)
	}
	return hints, nil
}

func (h *Hint) ReadByAnswer() ([]Hint, error) {
	rows, err := db.Query("SELECT id, question, answer, user_id FROM hints WHERE answer = ?", h.Answer)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var hints []Hint
	for rows.Next() {
		hint := Hint{}
		if err := rows.Scan(&hint.ID, &hint.Question, &hint.Answer, &hint.UserID); err != nil {
			return nil, err
		}
		hints = append(hints, hint)
	}
	return hints, nil
}

func (h *Hint) ReadByUserID() error {
	return db.QueryRow("SELECT id, question, answer FROM hints WHERE user_id = ?", h.UserID).
		Scan(&h.ID, &h.Question, &h.Answer)
}

func (h *Hint) ReadAll() ([]Hint, error) {
	rows, err := db.Query("SELECT id, question, answer, user_id FROM hints")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var hints []Hint
	for rows.Next() {
		hint := Hint{}
		if err := rows.Scan(&hint.ID, &hint.Question, &hint.Answer, &hint.UserID); err != nil {
			return nil, err
		}
		hints = append(hints, hint)
	}
	return hints, nil
}

func (h *Hint) Update() error {
	_, err := db.Exec("UPDATE hints SET question = ?, answer = ?, user_id = ? WHERE id = ?",
		h.Question, h.Answer, h.UserID, h.ID)
	if err != nil {
		return err
	}
	return nil
}

func (h *Hint) Delete() error {
	_, err := db.Exec("DELETE FROM hints WHERE id = ?", h.ID)
	if err != nil {
		return err
	}
	return nil
}

func (h *Hint) DeleteAll() error {
	_, err := db.Exec("DELETE FROM hints")
	if err != nil {
		return err
	}
	return nil
}
