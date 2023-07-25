package handlers

import (
	"TheFirst/config"
	"TheFirst/models"
	"errors"
	"fmt"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var logger *log.Logger
var localizer *i18n.Localizer

func init() {
	// 获取本地化实例
	localizer = i18n.NewLocalizer(config.Viperconfig.LocaleBundle, config.Viperconfig.App.Language)
	file, err := os.OpenFile(config.Viperconfig.App.Log+"/chitchat.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)
}

// 检查用户是否登录并有会话
func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
	cookie, err := request.Cookie("_cookie")
	if err == nil {
		sess = models.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check_session(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}

// 生成 HTML 模板
func generateHTML(writer http.ResponseWriter, data interface{}, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("views/%s/%s.html", config.Viperconfig.App.Language, file))
	}
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("layout").Funcs(funcMap)
	templates := template.Must(t.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", data)
}

// version
func Version() string {
	return "0.1"
}

// 记录日志信息
func info(args ...interface{}) {
	logger.SetPrefix("INFO ")
	logger.Println(args...)
}

func danger(args ...interface{}) {
	logger.SetPrefix("ERROR ")
	logger.Println(args...)
}

func warning(args ...interface{}) {
	logger.SetPrefix("WARNING ")
	logger.Println(args...)
}

// 异常处理统一重定向到错误页面
func errorMessage(writer http.ResponseWriter, request *http.Request, msg string) {
	url := []string{"/err?msg=", msg}
	http.Redirect(writer, request, strings.Join(url, ""), 302)
}

// 日期格式化
func formatDate(t time.Time) string {
	datetime := "2006-01-02 15:04:05"
	return t.Format(datetime)
}
