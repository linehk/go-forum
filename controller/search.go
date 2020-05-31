package controller

import (
	"net/http"
)

// search 渲染搜索页面。
func search(w http.ResponseWriter, r *http.Request) {
	if !logged(w, r) {
		msg(w, r, "未登录无法搜索")
		return
	}
	html(w, r, nil, "layout", "search")
}

// todo:fuzzySearchAction
func fuzzySearchAction(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		msgErr(w, r, "解析表单错误：", err)
		return
	}
}

// todo:exactSearchAction
func exactSearchAction(w http.ResponseWriter, r *http.Request) {
}
