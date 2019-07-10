package controller

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/linehk/go-forum/model"
)

func logged(w http.ResponseWriter, r *http.Request) bool {
	_, err := session(w, r)
	if err != nil {
		return false
	}
	return true
}

func session(w http.ResponseWriter, r *http.Request) (*model.Session, error) {
	c, err := r.Cookie("_cookie")
	if err != nil {
		return nil, err
	}
	s := &model.Session{UUID: c.Value}
	if err := s.ReadByUUID(); err != nil {
		return nil, err
	}
	return s, nil
}

func verify(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

// get
func signup(w http.ResponseWriter, r *http.Request) {
	html(w, r, nil, "layout", "public.navbar", "signup")
}

// get
func login(w http.ResponseWriter, r *http.Request) {
	html(w, r, nil, "layout", "public.navbar", "login")
}

// get
func logout(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("_cookie")
	if err != nil {
		msg(w, r, "get cookie err: ", err)
	}
	s := &model.Session{UUID: c.Value}
	if err := s.ReadByUUID(); err != nil {
		msg(w, r, "read by uuid err: ", err)
	}
	if err := s.Delete(); err != nil {
		msg(w, r, "delete session err: ", err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// post
func createUser(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		msg(w, r, "parse form err: ", err)
	}
	u := &model.User{
		Name:     r.PostFormValue("name"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
	}
	if err := u.Create(); err != nil {
		msg(w, r, "create user err: ", err)
	}
	http.Redirect(w, r, "/login", http.StatusFound)
}

// post
func auth(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		msg(w, r, "parse form err: ", err)
	}
	u := &model.User{Email: r.PostFormValue("email")}
	if err := u.ReadByEmail(); err != nil {
		msg(w, r, "read email err: ", err)
	}
	if !verify(r.PostFormValue("password"), u.Password) {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	s := &model.Session{UserID: u.ID}
	if err := s.Create(); err != nil {
		msg(w, r, "create session err: ", err)
	}
	if err := s.ReadByUserID(); err != nil {
		msg(w, r, "read by user_id err: ", err)
	}
	c := http.Cookie{
		Name:     "_cookie",
		Value:    s.UUID,
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusFound)
}
