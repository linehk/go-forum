package model

type User struct {
	ID       int
	Name     string
	Email    string
	Password string
}

func (u *User) Create() error {
	_, err := db.Exec("INSERT INTO users (name, email, password) VALUES (?, ?, ?)", u.Name, u.Email, hash(u.Password))
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ReadByID() error {
	return db.QueryRow("SELECT name, email, password FROM users WHERE id = ?", u.ID).
		Scan(&u.Name, &u.Email, &u.Password)
}

func (u *User) ReadByName() error {
	return db.QueryRow("SELECT id, email, password FROM users WHERE name = ?", u.Name).
		Scan(&u.ID, &u.Email, &u.Password)
}

func (u *User) ReadByEmail() error {
	return db.QueryRow("SELECT id, name, password FROM users WHERE email = ?", u.Email).
		Scan(&u.ID, &u.Name, &u.Password)
}

func (u *User) ReadAll() ([]User, error) {
	rows, err := db.Query("SELECT id, name, email, password FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) Update() error {
	_, err := db.Exec("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?",
		u.Name, u.Email, hash(u.Password), u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Delete() error {
	_, err := db.Exec("DELETE FROM users WHERE id = ?", u.ID)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) DeleteAll() error {
	_, err := db.Exec("DELETE FROM users")
	if err != nil {
		return err
	}
	return nil
}
