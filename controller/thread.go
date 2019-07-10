package controller

import (
	"net/http"

	"github.com/linehk/go-forum/model"
)

// threads 渲染新建主题时的页面。
func threads(w http.ResponseWriter, r *http.Request) {
	if logged(w, r) {
		html(w, r, nil, "layout", "private.navbar", "new.thread")
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// createThread 根据表单来建立主题。
func createThread(w http.ResponseWriter, r *http.Request) {
	s, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	if err := r.ParseForm(); err != nil {
		msg(w, r, "parse form err: ", err)
	}
	// 读取 session 对应的用户。
	u := &model.User{ID: s.UserID}
	if err := u.ReadByID(); err != nil {
		msg(w, r, "read user err: ", err)
	}
	// 为该用户建立主题。
	t := &model.Thread{Subject: r.PostFormValue("subject"), UserID: u.ID}
	if err := t.Create(); err != nil {
		msg(w, r, "create thread err: ", err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}

// 临时用于传入渲染的数据
type ReadView struct {
	Thread model.Thread
	Posts  []model.Post
}

// readThread 渲染 /threads/read?uuid= 页面。
func readThread(w http.ResponseWriter, r *http.Request) {
	uuid := r.URL.Query().Get("uuid")
	data := &ReadView{Thread: model.Thread{UUID: uuid}}
	if err := data.Thread.ReadByUUID(); err != nil {
		msg(w, r, "read thread err: ", err)
	}
	p := &model.Post{ThreadID: data.Thread.ID}
	posts, err := p.ReadByThreadID()
	if err != nil {
		msg(w, r, "read by thread_id err: ", err)
	}
	data.Posts = posts
	if logged(w, r) {
		html(w, r, data, "layout", "private.navbar", "private.thread")
	} else {
		html(w, r, data, "layout", "public.navbar", "public.thread")
	}
}

// createPost 用于建立主题对应的帖子。
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
		Content:  r.PostFormValue("content"),
		UserID:   s.UserID,
		ThreadID: t.ID,
	}
	if err := p.Create(); err != nil {
		msg(w, r, "create post err: ", err)
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
