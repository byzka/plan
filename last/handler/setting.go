package handler

import (
	"github.com/gin-gonic/gin"
	"last/model"
	"log"
	"net/http"
)

func Creat_s(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		log.Panicln("err", err.Error())
		context.String(http.StatusBadRequest, "wrong")
	}
	user.Creat_setting(user.Password)
	if user.Set.Favourite_local != user.Set.Enviroment {
		log.Panicln("err")
		return
	}
	context.Redirect(http.StatusMovedPermanently, "/setting/show")
}
func Show(context *gin.Context) {
	var u model.UserModel
	if err := context.Bind(&u); err != nil {
		log.Panicln("err", err)
	}
	answer, err := u.CHeck_Bid(u.Password)
	if err != nil {
		log.Panicln(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": answer.Set,
	})
}

func Change(context *gin.Context) {
	var user model.UserModel
	if err := context.Bind(&user); err != nil {
		log.Panicln()
	}
	answer, err := user.CHeck_Bid(user.Password)
	if err != nil {
		log.Panicln(err)
	}
	answer.Set.Enviroment = user.Set.Enviroment
	answer.Set.Favourite_local = user.Set.Favourite_local
	err = user.Update_set(user.ID)
	if err != nil {
		log.Panicln(err)
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": answer.Set,
	})
}
