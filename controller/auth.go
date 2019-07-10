package controller

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/linehk/go-forum/model"
)

// logged 判断是否已经登录。
func logged(w http.ResponseWriter, r *http.Request) bool {
	_, err := session(w, r)
	return err == nil
}

// session 根据 cookie 读取 session。
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

// verify 验证密码是否正确。
func verify(password, hash string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)); err != nil {
		return false
	}
	return true
}

// signup 渲染 /signup 页面。
func signup(w http.ResponseWriter, r *http.Request) {
	html(w, r, nil, "layout", "public.navbar", "signup")
}

// login 渲染 /login 页面。
func login(w http.ResponseWriter, r *http.Request) {
	html(w, r, nil, "layout", "public.navbar", "login")
}

// logout 表示登出，删除 cookie 对应的 session。
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

// createUser 根据表单里的内容建立用户。
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

// auth 判断表单中的密码是否正确，并建立 session 和 设置 cookie。
func auth(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		msg(w, r, "parse form err: ", err)
	}
	u := &model.User{Email: r.PostFormValue("email")}
	if err := u.ReadByEmail(); err != nil {
		msg(w, r, "read email err: ", err)
	}
	// 验证
	if !verify(r.PostFormValue("password"), u.Password) {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	// 建立 session
	s := &model.Session{UserID: u.ID}
	if err := s.Create(); err != nil {
		msg(w, r, "create session err: ", err)
	}
	// 读取 session
	if err := s.ReadByUserID(); err != nil {
		msg(w, r, "read by user_id err: ", err)
	}
	// 设置 cookie
	c := http.Cookie{
		Name:     "_cookie",
		Value:    s.UUID,
		HttpOnly: true,
	}
	http.SetCookie(w, &c)
	http.Redirect(w, r, "/", http.StatusFound)
}
