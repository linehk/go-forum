package controller

import (
	"net/http"

	"github.com/linehk/go-forum/model"
)

// get
func threads(w http.ResponseWriter, r *http.Request) {
	if logged(w, r) {
		html(w, r, nil, "layout", "private.navbar", "new.thread")
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// post
func createThread(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	if err := r.ParseForm(); err != nil {
		msg(w, r, "parse form err: ", err)
	}
	u := &model.User{ID: s.UserID}
	if err := u.ReadByID(); err != nil {
		msg(w, r, "read user err: ", err)
	}
	t := &model.Thread{Subject: r.PostFormValue("subject"), UserID: u.ID}
	if err := t.Create(); err != nil {
		msg(w, r, "create thread err: ", err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// get
func readThread(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")
	t := &model.Thread{UUID: uuid}
	if err := t.ReadByUUID(); err != nil {
		msg(w, r, "read thread err: ", err)
	}
	if logged(w, r) {
		html(w, r, t, "layout", "private.navbar", "private.thread")
	} else {
		html(w, r, t, "layout", "public.navbar", "public.thread")
	}
}

// post
func createPost(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	if err := r.ParseForm(); err != nil {
		msg(w, r, "parse form err: ", err)
	}
	t := &model.Thread{UUID: r.PostFormValue("uuid")}
	if err := t.ReadByUUID(); err != nil {
		msg(w, r, "read thread err: ", err)
	}
	p := &model.Post{
		Title:    r.PostFormValue("title"),
		Content:  r.PostFormValue("content"),
		UserID:   s.UserID,
		ThreadID: t.ID,
	}
	if err := p.Create(); err != nil {
		msg(w, r, "create post err: ", err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
