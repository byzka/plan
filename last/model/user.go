package model

import (
	"database/sql"
	"last/init_DB"
	"log"
)

type UserModel struct {
	Email        string `form:"email"`
	Password     string `form:"password"`
	ID           int    `form:"id"`
	Phone_number int    `form:"number"`
	Head         sql.NullString
	Set          Setting
}

func (user *UserModel) Save() int64 {
	result, err := init_DB.DB.Exec("insert into user.user (email, password) values (?,?);", user.Email, user.Password)
	if err != nil {
		log.Panicln("err", err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("err", err.Error())
	}
	return id
}

func (user *UserModel) Q_BY_email(email string) (UserModel, error) {
	u := UserModel{}
	row := init_DB.DB.QueryRow("select * from user where email = ?;", email)
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Phone_number, &u.Head)
	if err != nil {
		log.Panicln(err)
	}
	return u, err
}

func (user *UserModel) Q_BY_id(id int) (UserModel, error) {
	u := UserModel{}
	row := init_DB.DB.QueryRow("select * from user where id = ?;", id)
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Phone_number, &u.Head)
	if err != nil {
		log.Panicln(err)
	}
	return u, err
}

func (user *UserModel) Q_BY_PHone(phone int) (UserModel, error) {
	u := UserModel{}
	row := init_DB.DB.QueryRow("select * from user where phone_number = ?;", phone)
	err := row.Scan(&u.ID, &u.Email, &u.Password, &u.Phone_number, &u.Head)
	if err != nil {
		log.Panicln(err)
	}
	return u, err
}

func (user *UserModel) Update(id int) error {
	var stmt, e = init_DB.DB.Prepare("update user set password=?,head=?  where id=? ")
	if e != nil {
		log.Panicln("发生了错误", e.Error())
	}
	_, e = stmt.Exec(user.Password, user.Head.String, user.ID)
	if e != nil {
		log.Panicln("错误 e", e.Error())
	}

	return e
}
