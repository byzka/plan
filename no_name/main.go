package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/http"
	"no_name/model"
)

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	db.AutoMigrate()

	server := gin.Default()

	server.LoadHTMLGlob("templates/*")
	server.Static("/static", "./static")
	server.GET("/index", func(context *gin.Context) {
		context.HTML(http.StatusOK, "useradd.html", "")
	})
	//注册
	server.POST("/user/add", func(context *gin.Context) {
		u := model.User{State: model.Online}
		context.BindJSON(&u)
		//查询是否存在
		res := db.Where("name = ? ", u.Name).First(&model.User{})
		if res.RowsAffected != 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg": "注册失败，已存在用户",
			})
		} else {
			db.Create(&u)
			context.JSON(http.StatusOK, gin.H{
				"msg": "注册成功",
			})
		}
	})
	//登陆
	server.POST("/user/login", func(context *gin.Context) {
		u := model.User{State: model.Online}
		context.BindJSON(&u)
		res := db.Where("name = ? AND passwd = ? ", u.Name, u.Passwd).First(&model.User{})
		if res.RowsAffected != 0 {
			context.JSON(http.StatusOK, gin.H{
				"msg": "登陆成功",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "登陆失败",
			})
		}
	})

	server.POST("/user/logout", func(context *gin.Context) {
		u := model.User{}
		context.BindJSON(&u)
		res := db.Where("name = ? ", u.Name).First(&model.User{})
		if res.RowsAffected != 0 {
			var user model.User
			db.Where("name = ? AND passwd = ? ", u.Name, u.Passwd).Take(&user)
			user.State = model.Notonline
			db.Save(&user)
			context.JSON(http.StatusOK, gin.H{
				"msg": "登出成功",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "username or passwd wrong",
			})
		}

	})
	//更改信息
	server.POST("/user/change", func(context *gin.Context) {
		u := model.User{}
		context.BindJSON(&u)
		res := db.Where("name = ? AND passwd = ? ", u.Name, u.Passwd).First(&model.User{})
		if res.RowsAffected != 0 {
			var user model.User
			db.Where("name = ? ", u.Name).Take(&user)
			//change
			user.Name = u.Name
			user.Passwd = u.Passwd
			user.Email = u.Email
			user.Birthday = u.Birthday
			user.Phone_number = u.Phone_number
			db.Save(&user)
			context.JSON(http.StatusOK, gin.H{
				"msg": "更改成功",
			})
		} else {
			context.JSON(http.StatusOK, gin.H{
				"msg": "信息错误",
			})
		}
	})

	server.Run(":9999")
}
