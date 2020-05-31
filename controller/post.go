package controller

import (
	"net/http"

	"github.com/linehk/go-forum/model"
)

// createPostAction 用于建立主题对应的帖子。
func createPostAction(w http.ResponseWriter, r *http.Request) {
	if !logged(w, r) {
		msg(w, r, "未登录无法新建帖子")
		return
	}

	if err := r.ParseForm(); err != nil {
		msgErr(w, r, "解析表单错误：", err)
		return
	}

	t := &model.Thread{UUID: r.PostFormValue("uuid")}
	if err := t.ReadByUUID(); err != nil {
		msgErr(w, r, "通过 UUID 读取主题错误：", err)
		return
	}

	s, err := session(w, r)
	if err != nil {
		msgErr(w, r, "读取会话错误：", err)
		return
	}

	p := &model.Post{
		Content:  r.PostFormValue("content"),
		UserID:   s.UserID,
		ThreadID: t.ID,
	}
	if err := p.Create(); err != nil {
		msgErr(w, r, "创建帖子错误：", err)
		return
	}

	http.Redirect(w, r, "/threads/read?uuid="+t.UUID, http.StatusFound)
}

// todo:updatePost
func updatePost(w http.ResponseWriter, r *http.Request) {
}

// todo:updatePostAction
func updatePostAction(w http.ResponseWriter, r *http.Request) {
}

// todo:deletePostAction
func deletePostAction(w http.ResponseWriter, r *http.Request) {
}
