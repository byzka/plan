package handler

import (
	"github.com/gin-gonic/gin"
	"last/model"
	"net/http"
)

func S_Q(context *gin.Context) {
	context.HTML(http.StatusOK, "search.html", gin.H{
		"title": "input your search",
	})
	var s model.Search
	if err := context.ShouldBind(&s); err != nil {
		panic("wrong")
	}
	if s.Type != "" {
		context.SetCookie("cookie_type", s.Type, 1000, "/search/a", "localhost", false, true)
	} else if s.Name != "" {
		context.SetCookie("cookie_name", s.Name, 1000, "/search/a", "localhost", false, true)
	} else if s.Local != "" {
		context.SetCookie("cookie_local", s.Local, 1000, "/search/a", "localhost", false, true)
	} else {
		context.SetCookie("cookie_all", "", 1000, "/search/a", "localhost", false, true)
	}
	context.Redirect(http.StatusMovedPermanently, "/search/a")
}

func S_A(context *gin.Context) {
	var s model.Search
	var err error
	s.Type, err = context.Cookie("cookie_type")
	if err != nil {
		panic("wrong")
	}
	s.Name, err = context.Cookie("cookie_name")
	if err != nil {
		panic("wrong")
	}
	s.Local, err = context.Cookie("cookie_local")
	if err != nil {
		panic("wrong")
	}
	answer, err := s.Search_sth()
	context.JSON(http.StatusOK, gin.H{
		"you got this": answer,
	})
}
