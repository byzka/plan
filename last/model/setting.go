package model

import (
	"last/init_DB"
	"log"
)

type Setting struct {
	Favourite_local string
	Enviroment      string
}

func (u *UserModel) Creat_setting(ps string) {
	if u.Password != ps {
		panic("no same")
		return
	}
	u.Set.Favourite_local = "all"
	u.Set.Enviroment = "all"
}

func (u *UserModel) CHeck_Bid(ps string) (UserModel, error) {
	user := UserModel{}
	row := init_DB.DB.QueryRow("select * from user where id = ?;", u.ID)
	err := row.Scan(&user.ID, &user.Email, &user.Password, &user.Phone_number, &user.Head, &user.Set)
	if err != nil {
		log.Panicln(err)
	}
	if ps != u.Password {
		log.Panicln(err)
	}
	return user, err
}

func (user *UserModel) Update_set(id int) error {
	var stmt, e = init_DB.DB.Prepare("update user set local=?,enviroment=?  where id=? ")
	if e != nil {
		log.Panicln("发生了错误", e.Error())
	}
	_, e = stmt.Exec(user.Set.Favourite_local, user.Set.Enviroment, user.ID)
	if e != nil {
		log.Panicln("错误 e", e.Error())
	}

	return e
}
