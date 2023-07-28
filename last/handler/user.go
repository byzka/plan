package handler

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"last/model"
	"last/utils"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "save successful,welcome:"+username)
}

func UserSaveQ(context *gin.Context) {
	username := context.Query("name")
	age := context.Query("age")
	context.String(http.StatusOK, "username:"+username+"age:"+age)
}

func User_Register(context *gin.Context) {

	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Println("err->", err.Error())
		context.String(http.StatusBadRequest, "wrong")
	}
	id := user.Save()
	log.Println("id is :", id)
	context.Redirect(http.StatusMovedPermanently, "/")
}
func User_Login(context *gin.Context) {
	var user model.UserModel
	if err := context.Bind(&user); err != nil {
		log.Panicln(err)
	}
	u, err := user.Q_BY_email(user.Email)
	if err != nil {
		log.Panicln(err)
	}
	if u.Password == user.Password {
		context.SetCookie("user_cookie", string(u.ID), 1000, "/", "/localhost", false,
			true,
		)
		log.Println("success,", u.Email)
		context.HTML(http.StatusOK, "index.html", gin.H{
			"email": u.Email,
		})
	}
}

func UserProfile(context *gin.Context) {
	id := context.Query("id")
	var user model.UserModel
	id1, err := strconv.Atoi(id)
	us, err1 := user.Q_BY_id(id1)
	if err != nil || err1 != nil {
		context.HTML(http.StatusOK, "error.html", gin.H{
			"error": err,
		})
	}
	context.HTML(http.StatusOK, "user_profile.html", gin.H{
		"user": us,
	})
}

func Update_UserProfile(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		context.HTML(http.StatusOK, "error.html", gin.H{
			"error": err,
		})
		log.Panicln("绑定错误", err.Error())
	}
	file, err := context.FormFile("head-file")
	if err != nil {
		context.HTML(http.StatusOK, "error.html", gin.H{
			"error": err,
		})
		log.Panicln("上传错误", err.Error())
	}
	path := utils.Root_Path()
	path = path + "head\\"
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("无法创建文件夹", err.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, path+fileName)
	if err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("无法保存文件", err.Error())
	}
	headUrl := "http://localhost:8080/head/" + fileName
	user.Head = sql.NullString{String: headUrl}
	err = user.Update(user.ID)
	if err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err,
		})
		log.Panicln("数据无法更新", err.Error())
	}
	context.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(int(user.ID)))
}

func Logout(context *gin.Context) {
	var user model.UserModel
	if err := context.Bind(&user); err != nil {
		log.Panicln(err)
	}
	context.SetCookie("user_cookie", string(user.ID), -1, "/", "/localhost", false,
		true,
	)
	context.JSON(200, gin.H{
		"msg": "quit success",
	})
	context.Redirect(http.StatusMovedPermanently, "/")
}
