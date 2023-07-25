package handlers

import (
	"TheFirst/models"
	"fmt"
	"net/http"
)

// GET /login
// 登录页面
func Login(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "login")
}

// GET /signup
// 注册页面
func Signup(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, nil, "auth.layout", "navbar", "signup")
}

// POST /signup
// 注册新用户
func SignupAccount(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "Cannot parse form")
	}
	user := models.User{
		Name:     request.PostFormValue("name"),
		Email:    request.PostFormValue("email"),
		Password: request.PostFormValue("password"),
	}
	if err := user.Creat(); err != nil {
		danger(err, "Cannot create user")
	}
	http.Redirect(writer, request, "/login", 302)
}

// POST /authenticate
// 通过邮箱和密码字段对用户进行认证
func Authenticate(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	user, err := models.UserByEmail(request.PostFormValue("email"))
	if err != nil {
		danger(err, "Cannot find user")
	}
	if user.Password == models.Encrypt(request.PostFormValue("password")) {
		session, err := user.CreatSession()
		if err != nil {
			danger(err, "Cannot create session")
		}
		cookie := http.Cookie{
			Name:     "_cookie",
			Value:    session.Uuid,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
		http.Redirect(writer, request, "/", 302)
	} else {
		http.Redirect(writer, request, "/login", 302)
	}
}

// GET /logout
// 用户退出
func Logout(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_cookie")
	if err != http.ErrNoCookie {
		warning(err, "Failed to get cookie")
		session := models.Session{Uuid: cookie.Value}
		session.DeleteByUUID()
	}
	http.Redirect(writer, request, "/", 302)
}

func Search_sth(writer http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		danger(err, "cannot search")
	}
	sword := models.Answer{
		Type:  request.PostFormValue("type"),
		Name:  request.PostFormValue("name"),
		Local: request.PostFormValue("local"),
	}
	if sword.Name != "" {
		cookie := http.Cookie{
			Name:     "_search",
			Value:    sword.Name,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
	}
	if sword.Local != "" {
		cookie := http.Cookie{
			Name:     "_search",
			Value:    sword.Local,
			HttpOnly: true,
		}
		http.SetCookie(writer, &cookie)
	}
	http.Redirect(writer, request, "/search/success", 302)
}

func Search_success(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("_search")
	if err != nil {
		return
	}
	sth := cookie.Value
	a, err := models.Find_name(sth)
	fmt.Println(a)
}
