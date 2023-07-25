package models

import (
	"TheFirst/config"
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
)

var DB *sql.DB

func init() {
	var err error
	driver := config.Viperconfig.Db.Driver
	source := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=true", config.Viperconfig.Db.User, config.Viperconfig.Db.Password,
		config.Viperconfig.Db.Address, config.Viperconfig.Db.Database)
	DB, err = sql.Open(driver, source)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func CreateUuID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}
	//ID处理
	u[8] = (u[8] | 0x40) & 0x7F
	//最后四位变为版本号
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

// 使用SHA-1哈希铭文
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}
