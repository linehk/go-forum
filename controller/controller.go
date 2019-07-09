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

// get
func index(w http.ResponseWriter, r *http.Request) {
	threads, err := model.Thread{}.ReadAll()
	if err != nil {
		msg(w, r, "read threads err: ", err)
	}
	if logged(w, r) {
		html(w, r, threads, "layout", "private.navbar", "index")
	} else {
		html(w, r, threads, "layout", "public.navbar", "index")
	}
}

// get
func err(w http.ResponseWriter, r *http.Request) {
	msg := r.URL.Query().Get("msg")
	if logged(w, r) {
		html(w, r, msg, "layout", "private.navbar", "error")
	} else {
		html(w, r, msg, "layout", "public.navbar", "error")
	}
}

func msg(w http.ResponseWriter, r *http.Request, msg string, err error) {
	http.Redirect(w, r, "/err?msg="+msg+err.Error(), http.StatusFound)
}

func parse(names ...string) *template.Template {
	var files []string
	for _, f := range names {
		files = append(files, fmt.Sprintf("view/%s.html", f))
	}
	return template.Must(template.New("layout").ParseFiles(files...))
}

func html(w http.ResponseWriter, r *http.Request, data interface{}, names ...string) {
	var files []string
	for _, f := range names {
		files = append(files, fmt.Sprintf("view/%s.html", f))
	}
	if err := template.Must(template.ParseFiles(files...)).ExecuteTemplate(w, "layout", data); err != nil {
		msg(w, r, "execute template err: ", err)
	}
}
