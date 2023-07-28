package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func Index(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"title": strings.ToLower(context.Request.Method) + "method",
	})
}
