package controller

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/linehk/go-forum/model"
)

func Setup() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/err", err)

	// auth
	mux.HandleFunc("/signup", signup)
	mux.HandleFunc("/signup_action", signupAction)
	mux.HandleFunc("/login", login)
	mux.HandleFunc("/login_action", loginAction)
	mux.HandleFunc("/reset", reset)
	mux.HandleFunc("/reset_action", resetAction)
	mux.HandleFunc("/set_admin", setAdmin)
	mux.HandleFunc("/set_admin_action", setAdminAction)
	mux.HandleFunc("/logout_action", logoutAction)

	// thread
	mux.HandleFunc("/threads/create", createThread)
	mux.HandleFunc("/threads/create_action", createThreadAction)
	mux.HandleFunc("/threads/read", readThread)
	mux.HandleFunc("/threads/update", updateThread)
	mux.HandleFunc("/threads/update_action", updateThreadAction)
	mux.HandleFunc("/threads/delete_action", deleteThreadAction)

	// post
	mux.HandleFunc("/posts/create_action", createPostAction)
	mux.HandleFunc("/posts/update", updatePost)
	mux.HandleFunc("/posts/update_action", updatePostAction)
	mux.HandleFunc("/posts/delete_action", deletePostAction)

	// search
	mux.HandleFunc("/search", search)
	mux.HandleFunc("/fuzzy_search_action", fuzzySearchAction)
	mux.HandleFunc("/exact_search_action", exactSearchAction)

	return mux
}

// 临时用于传入渲染的数据
type Index struct {
	ThreadSubject   string
	ThreadCreatedAt time.Time
	ThreadUUID      string
	Username        string
	PostCount       int
}

// index 获取所有主题，并渲染 /index 页面。
func index(w http.ResponseWriter, r *http.Request) {
	if !logged(w, r) {
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}

	// 填充数据
	var data []Index

	t := &model.Thread{}
	threads, err := t.ReadAll()
	if err != nil {
		msgErr(w, r, "读取全部主题错误：", err)
		return
	}
	for _, t := range threads {
		var i Index
		i.ThreadSubject = t.Subject
		i.ThreadCreatedAt = t.CreatedAt
		i.ThreadUUID = t.UUID

		u := &model.User{ID: t.UserID}
		if err := u.ReadByID(); err != nil {
			msgErr(w, r, "通过 ID 读取用户错误：", err)
			return
		}
		i.Username = u.Username

		p := &model.Post{ThreadID: t.ID}
		c, err := p.CountByThreadID()
		if err != nil {
			msgErr(w, r, "通过主题 ID 统计帖子错误：", err)
			return
		}
		i.PostCount = c

		data = append(data, i)
	}

	html(w, r, data, "layout", "index")
}

// err 根据传入的 msg URL 参数来渲染 /err 错误页面。
func err(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	html(w, r, msg, "layout", "err")
}

// msgErr 组合传入的参数和错误信息重定向到 /err?msg= 错误页面。
func msgErr(w http.ResponseWriter, r *http.Request, msg string, err error) {
	http.Redirect(w, r, "/err?msg="+msg+err.Error(), http.StatusFound)
}

// msg 组合传入的参数重定向到 /err?msg= 错误页面。
func msg(w http.ResponseWriter, r *http.Request, msg string) {
	http.Redirect(w, r, "/err?msg="+msg, http.StatusFound)
}

// html 根据传入的数据 data 和应该渲染的文件 names 来渲染页面。
func html(w http.ResponseWriter, r *http.Request, data interface{}, names ...string) {
	var files []string
	for _, f := range names {
		files = append(files, fmt.Sprintf("view/%s.html", f))
	}
	if err := template.Must(template.ParseFiles(files...)).ExecuteTemplate(w, "layout", data); err != nil {
		msgErr(w, r, "渲染模板错误：", err)
	}
}
