package controller

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"github.com/linehk/go-forum/model"
)

// signup 渲染注册页面。
func signup(w http.ResponseWriter, r *http.Request) {
	html(w, r, nil, "layout", "signup")
}

// signupAction 根据表单里的内容注册。
// 全部表单内容为：map[answer:[dota2] confirmPassword:[line] email:[linehk@gmail.com] hint:[1] password:[line] username:[linehk]]
func signupAction(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		msgErr(w, r, "解析表单错误：", err)
		return
	}

	// 确认密码
	if r.PostFormValue("password") != r.PostFormValue("confirmPassword") {
		msg(w, r, "两次密码不一致")
		return
	}

	// 读取表单内容，并创建用户
	u := &model.User{
		Username: r.PostFormValue("username"),
		Email:    r.PostFormValue("email"),
		Password: r.PostFormValue("password"),
		Admin:    r.PostFormValue("admin"),
	}
	if err := u.Create(); err != nil {
		msgErr(w, r, "创建用户错误：", err)
		return
	}

	// 根据 username 读取
	if err := u.ReadByUsername(); err != nil {
		msgErr(w, r, "通过用户名读取用户错误：", err)
		return
	}

	// 创建安全提问，不填写安全提问时，表单 hint 的内容为 ignore，answer 为空
	h := &model.Hint{
		Question: r.PostFormValue("hint"),
		Answer:   r.PostFormValue("answer"),
		UserID:   u.ID,
	}
	if err := h.Create(); err != nil {
		msgErr(w, r, "创建安全提问错误：", err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

// login 渲染登录页面。
func login(w http.ResponseWriter, r *http.Request) {
	html(w, r, nil, "layout", "login")
}

// loginAction 判断表单中的密码是否正确，并建立 session 和 设置 cookie。
// 全部表单内容为：map[admin:[yes] answer:[dota2] autoLogin:[yes] email:[sulinehk@gmail.com] hint:[1] password:[admin]]
// 部分表单内容为：map[answer:[] email:[sulinehk@gmail.com] hint:[ignore] password:[admin]]
func loginAction(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		msgErr(w, r, "解析表单错误：", err)
		return
	}

	// 根据 email 读取用户
	u := &model.User{Email: r.PostFormValue("email")}
	if err := u.ReadByEmail(); err != nil {
		msgErr(w, r, "通过邮箱读取用户错误：", err)
		return
	}

	// 验证密码
	if !verify(r.PostFormValue("password"), u.Password) {
		msg(w, r, "密码不正确")
		return
	}

	// 验证管理员
	if r.PostFormValue("admin") != u.Admin {
		msg(w, r, "管理员信息不正确")
		return
	}

	// 验证安全提问
	h := &model.Hint{UserID: u.ID}
	if err := h.ReadByUserID(); err != nil {
		msgErr(w, r, "通过用户 ID 读取安全提问错误：", err)
		return
	}
	if h.Question != r.PostFormValue("hint") {
		msg(w, r, "安全提问不正确")
		return
	}
	if h.Answer != r.PostFormValue("answer") {
		msg(w, r, "安全提问不正确")
		return
	}

	// 建立 session
	s := &model.Session{UserID: u.ID}
	if err := s.Create(); err != nil {
		msgErr(w, r, "创建会话错误：", err)
		return
	}

	// 读取 session
	if err := s.ReadByUserID(); err != nil {
		msgErr(w, r, "通过用户 ID 读取会话错误", err)
		return
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

// reset 渲染重置密码页面。
func reset(w http.ResponseWriter, r *http.Request) {
	html(w, r, nil, "layout", "reset")
}

// resetAction 根据表单内容重置密码。
func resetAction(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		msgErr(w, r, "解析表单错误：", err)
		return
	}

	// 根据 email 读取 user
	u := &model.User{Email: r.PostFormValue("email")}
	if err := u.ReadByEmail(); err != nil {
		msgErr(w, r, "通过邮箱读取用户错误：", err)
		return
	}

	// 验证密码
	if !verify(r.PostFormValue("password"), u.Password) {
		msg(w, r, "密码不正确")
		return
	}

	// 验证安全提问
	h := &model.Hint{UserID: u.ID}
	if err := h.ReadByUserID(); err != nil {
		msgErr(w, r, "通过用户 ID 读取安全提问错误：", err)
		return
	}
	if h.Question != r.PostFormValue("hint") && h.Answer != r.PostFormValue("answer") {
		msg(w, r, "安全提问不正确")
		return
	}

	// 修改密码
	u.Password = r.PostFormValue("newPassword")
	if err := u.Update(); err != nil {
		msgErr(w, r, "更新用户错误：", err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

// setAdmin 根据是否为管理员，渲染设置管理员页面。
func setAdmin(w http.ResponseWriter, r *http.Request) {
	if !logged(w, r) {
		msg(w, r, "未登录无法设置管理员")
		return
	}

	// 读取 session
	s, err := session(w, r)
	if err != nil {
		msgErr(w, r, "读取会话错误：", err)
		return
	}

	// 根据 session 里的 user_id 读取 user
	u := &model.User{ID: s.UserID}
	if err := u.ReadByID(); err != nil {
		msgErr(w, r, "通过 ID 读取用户错误：", err)
		return
	}

	// 判断是否为管理员
	if u.Admin != "yes" {
		msg(w, r, "非管理员不允许设置管理员")
		return
	}

	html(w, r, nil, "layout", "set_admin")
}

// setAdminAction 设置管理员。
func setAdminAction(w http.ResponseWriter, r *http.Request) {
	if !logged(w, r) {
		msg(w, r, "未登录无法设置管理员")
		return
	}

	if err := r.ParseForm(); err != nil {
		msgErr(w, r, "解析表单错误：", err)
		return
	}

	// 用户名确认
	if r.PostFormValue("username") != r.PostFormValue("confirmUsername") {
		msg(w, r, "用户名不一致")
		return
	}

	// 根据 username 读取 user
	u := &model.User{Username: r.PostFormValue("username")}
	if err := u.ReadByUsername(); err != nil {
		msgErr(w, r, "通过用户名读取用户错误：", err)
		return
	}

	// 更新管理员为 yes
	u.Admin = "yes"
	if err := u.Update(); err != nil {
		msgErr(w, r, "更新用户错误：", err)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}

// logoutAction 表示登出，删除 cookie 对应的 session。
func logoutAction(w http.ResponseWriter, r *http.Request) {
	// 读取 session
	s, err := session(w, r)
	if err != nil {
		msgErr(w, r, "读取会话错误：", err)
		return
	}

	// 删除 session
	if err := s.Delete(); err != nil {
		msgErr(w, r, "删除会话错误：", err)
		return
	}

	http.Redirect(w, r, "/login", http.StatusFound)
}

// logged 判断是否已经登录。
func logged(w http.ResponseWriter, r *http.Request) bool {
	_, err := session(w, r)
	return err == nil
}

// session 根据 cookie 读取 session。
func session(w http.ResponseWriter, r *http.Request) (*model.Session, error) {
	// 读取 cookie，值为 session 的 uuid
	c, err := r.Cookie("_cookie")
	if err != nil {
		return nil, err
	}

	// 根据 uuid 读取 session
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
