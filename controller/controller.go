package controller

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/linehk/go-forum/model"
)

func Setup() *http.ServeMux {
	mux := http.NewServeMux()

	// get
	mux.HandleFunc("/", index)
	// get
	mux.HandleFunc("/err", err)

	// get
	mux.HandleFunc("/signup", signup)
	// get
	mux.HandleFunc("/login", login)
	// get
	mux.HandleFunc("/logout", logout)

	// post
	mux.HandleFunc("/users", createUser)
	// post
	mux.HandleFunc("/auth", auth)

	// get
	mux.HandleFunc("/threads/create", threads)
	// post
	mux.HandleFunc("/threads", createThread)

	// get
	mux.HandleFunc("/threads/read", readThread)
	// post
	mux.HandleFunc("/posts", createPost)

	return mux
}

// index 获取所有主题，并渲染 /index 页面。
func index(w http.ResponseWriter, r *http.Request) {
	t := &model.Thread{}
	threads, err := t.ReadAll()
	if err != nil {
		msg(w, r, "read threads err: ", err)
	}
	if logged(w, r) {
		html(w, r, threads, "layout", "private.navbar", "index")
	} else {
		html(w, r, threads, "layout", "public.navbar", "index")
	}
}

// err 根据传入的 msg URL 参数来渲染 /err 错误页面。
func err(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if logged(w, r) {
		html(w, r, msg, "layout", "private.navbar", "error")
	} else {
		html(w, r, msg, "layout", "public.navbar", "error")
	}
}

// msg 组合传入的参数重定向到 /err?msg= 错误页面。
func msg(w http.ResponseWriter, r *http.Request, msg string, err error) {
	http.Redirect(w, r, "/err?msg="+msg+err.Error(), http.StatusFound)
}

// html 根据传入的数据 data 和应该渲染的文件 names 来渲染页面。
func html(w http.ResponseWriter, r *http.Request, data interface{}, names ...string) {
	var files []string
	for _, f := range names {
		files = append(files, fmt.Sprintf("view/%s.html", f))
	}
	if err := template.Must(template.ParseFiles(files...)).ExecuteTemplate(w, "layout", data); err != nil {
		msg(w, r, "execute template err: ", err)
	}
}
