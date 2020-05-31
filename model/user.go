package model

type User struct {
	ID       int
	Username string
	Email    string
	Password string
	Admin    string
}

func (u *User) Create() error {
	_, err := db.Exec("INSERT INTO users (username, email, password, admin) VALUES (?, ?, ?, ?)",
		u.Username, u.Email, hash(u.Password), u.Admin)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ReadByID() error {
	return db.QueryRow("SELECT username, email, password, admin FROM users WHERE id = ?", u.ID).
		Scan(&u.Username, &u.Email, &u.Password, &u.Admin)
}

func (u *User) ReadByUsername() error {
	return db.QueryRow("SELECT id, email, password, admin FROM users WHERE username = ?", u.Username).
		Scan(&u.ID, &u.Email, &u.Password, &u.Admin)
}

func (u *User) ReadByEmail() error {
	return db.QueryRow("SELECT id, username, password, admin FROM users WHERE email = ?", u.Email).
		Scan(&u.ID, &u.Username, &u.Password, &u.Admin)
}

func (u *User) ReadByAdmin() ([]User, error) {
	rows, err := db.Query("SELECT id, username, email, password, admin FROM users WHERE admin = ?", u.Admin)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Admin); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) ReadAll() ([]User, error) {
	rows, err := db.Query("SELECT id, username, email, password, admin FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []User
	for rows.Next() {
		user := User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Admin); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u *User) Update() error {
	_, err := db.Exec("UPDATE users SET username = ?, email = ?, password = ?, admin = ? WHERE id = ?",
		u.Username, u.Email, hash(u.Password), u.Admin, u.ID)
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
