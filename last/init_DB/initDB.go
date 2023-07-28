package init_DB

import (
	"database/sql"
	"log"
)

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("mysql", "root:1234@tcp(127.0.0.1:3306)/ginhello")
	if err != nil {
		log.Panicln("err:", err.Error())
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(10)
}
