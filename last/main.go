package main

import (
	"github.com/gin-gonic/gin"
	"last/init_router"
	"net/http"
	"strings"
)

func main() {
	r := init_router.SetupRouter()

	r.Run()
}

func alluse(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
