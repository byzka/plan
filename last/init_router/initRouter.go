package init_router

import (
	"github.com/gin-gonic/gin"
	"last/handler"
	"last/utils"
	"net/http"
	"strings"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/statics", "./statics")
	r.StaticFile("/favicon.ico", "./favicon.ico")
	r.StaticFS("/avatar", http.Dir(utils.Root_Path()+"head/"))

	index := r.Group("/")
	{
		index.GET("", alluse)
		index.PUT("", alluse)
		index.POST("", alluse)
		index.DELETE("", alluse)
	}
	userRoute := r.Group("/user")
	{
		userRoute.GET("/:name", handler.UserSave)
		userRoute.GET("", handler.UserSaveQ)
		userRoute.POST("/register", handler.User_Register)
		userRoute.GET("/profile", handler.UserProfile)
		userRoute.POST("/login", handler.User_Login)
		userRoute.POST("/logout", handler.Logout)
		userRoute.POST("/change", handler.Update_UserProfile)

	}

	searchRoute := r.Group("/search")
	{
		searchRoute.GET("/a", handler.S_A)
		searchRoute.POST("/q", handler.S_Q)
	}
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "welcome")
	})
	return r
}
func alluse(context *gin.Context) {
	context.String(http.StatusOK, "hello gin "+strings.ToLower(context.Request.Method)+" method")
}
