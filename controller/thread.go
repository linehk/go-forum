package controller

import (
	"net/http"
	"time"

	"github.com/linehk/go-forum/model"
)

// createThread 渲染新建主题页面。
func createThread(w http.ResponseWriter, r *http.Request) {
	if !logged(w, r) {
		msg(w, r, "未登录无法新建主题")
		return
	}
	html(w, r, nil, "layout", "create_thread")
}

// createThreadAction 根据表单来建立主题。
func createThreadAction(w http.ResponseWriter, r *http.Request) {
	if !logged(w, r) {
		msg(w, r, "未登录无法新建主题")
		return
	}

	if err := r.ParseForm(); err != nil {
		msgErr(w, r, "解析表单错误：", err)
		return
	}

	// 读取 session。
	s, err := session(w, r)
	if err != nil {
		msgErr(w, r, "读取会话错误：", err)
		return
	}

	// 读取 session 对应的用户。
	u := &model.User{ID: s.UserID}
	if err := u.ReadByID(); err != nil {
		msgErr(w, r, "通过 ID 读取用户错误：", err)
		return
	}

	// 为该用户建立主题。
	t := &model.Thread{Subject: r.PostFormValue("subject"), UserID: u.ID}
	if err := t.Create(); err != nil {
		msgErr(w, r, "创建主题错误：", err)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// 临时用于传入渲染的数据
type ReadThread struct {
	ThreadSubject   string
	Username        string
	ThreadCreatedAt time.Time
	PostCount       int
	Posts           []model.Post
	ThreadUUID      string
}

// readThread 渲染 /threads/read?uuid= 页面。
func readThread(w http.ResponseWriter, r *http.Request) {
	if !logged(w, r) {
		msg(w, r, "未登录无法阅读主题")
		return
	}

	// 填充数据
	var data ReadThread

	t := &model.Thread{UUID: r.URL.Query().Get("uuid")}
	if err := t.ReadByUUID(); err != nil {
		msgErr(w, r, "通过 UUID 读取主题错误：", err)
		return
	}
	data.ThreadSubject = t.Subject
	data.ThreadCreatedAt = t.CreatedAt
	data.ThreadUUID = t.UUID

	u := &model.User{ID: t.UserID}
	if err := u.ReadByID(); err != nil {
		msgErr(w, r, "通过 ID 读取用户错误：", err)
		return
	}
	data.Username = u.Username

	p := &model.Post{ThreadID: t.ID}
	c, err := p.CountByThreadID()
	if err != nil {
		msgErr(w, r, "通过主题 ID 统计错误：", err)
		return
	}
	data.PostCount = c

	posts, err := p.ReadByThreadID()
	if err != nil {
		msgErr(w, r, "通过主题 ID 读取帖子错误：", err)
	}
	data.Posts = posts

	html(w, r, data, "layout", "read_thread")
}

// todo:updateThread
func updateThread(w http.ResponseWriter, r *http.Request) {
}

// todo:updateThreadAction
func updateThreadAction(w http.ResponseWriter, r *http.Request) {
}

// todo:deleteThreadAction
func deleteThreadAction(w http.ResponseWriter, r *http.Request) {
}
