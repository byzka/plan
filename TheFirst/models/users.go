package models

import (
	"time"
)

type User struct {
	ID       int
	UUID     int
	Name     string
	Age      int
	Email    string
	Password string
	CreatAT  time.Time
}

// 创建session
func (user *User) CreatSession() (session Session, err error) {
	statement := "insert into sessions (uuid, email, user_id, created_at) values (?, ?, ?, ?)"
	stmtin, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()
	uuid := CreateUuID()
	stmtin.Exec(uuid, user.Email, user.ID, time.Now())
	stmtout, err := DB.Prepare("select id, uuid, email, user_id, created_at from sessions where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// 获取已存在用户的Session
func (user *User) Session() (session Session, err error) {
	session = Session{}
	err = DB.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE user_id = ?", user.ID).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	return
}

// 创建一个新用户饼保存session
func (user *User) Creat() (err error) {
	statement := "insert into users (uuid, name, email, password, created_at) values (?, ?, ?, ?, ?)"
	stmtin, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()
	uuid := CreateUuID()
	stmtin.Exec(uuid, user.Name, user.Email, Encrypt(user.Password), time.Now())
	stmtout, err := DB.Prepare("select id, uuid, created_at from users where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()
	err = stmtout.QueryRow(uuid).Scan(&user.ID, &user.UUID, &user.CreatAT)
	return
}

// session里删除用户
func (user *User) Delete() (err error) {
	statement := "delete from users where id = ?"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.ID)
	return
}

// 更新用户信息
func (user *User) Update() (err error) {
	statement := "update users set name = ?, email = ? where id = ?"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Email, user.ID)
	return
}

// 数据库删除所有用户
func (user *User) DeleteAll() (err error) {
	statement := "delete from users"
	_, err = DB.Exec(statement)
	return
}

// 获取所有用户并返回
func (user *User) Users() (users []User, err error) {
	rows, err := DB.Query("SELECT id, uuid, name, email, password, created_at FROM users")
	if err != nil {
		return
	}
	for rows.Next() {
		u := User{}
		if err = rows.Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatAT); err != nil {
			return
		}
		users = append(users, u)
	}
	rows.Close()
	return
}

// 通过邮件获取用户
func UserByEmail(email string) (user User, err error) {
	user = User{}
	err = DB.QueryRow("SELECT id, uuid, name, email, password, created_at FROM users WHERE email = ?", email).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatAT)
	return
}

// 通过UUID获取用户
func UserByUUID(uuid string) (user User, err error) {
	user = User{}
	err = DB.QueryRow("SELECT id, uuid, name, email, password, creat_at FROM users WHERE uuid = ?", uuid).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.Password, &user.CreatAT)
	return
}

// 添加新线程
func (user *User) CreatThread(topic string) (conv Thread, err error) {
	statement := "insert into threads (uuid, topic, user_id, created_at) values (?, ?, ?, ?)"
	stmtin, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := CreateUuID()
	stmtin.Exec(uuid, topic, user.ID, time.Now())

	stmtout, err := DB.Prepare("select id, uuid, topic, user_id, created_at from threads where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	err = stmtout.QueryRow(uuid).Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// 为主题中创建帖子
func (user *User) CreatePost(conv Thread, body string) (post Post, err error) {
	statement := "insert into posts (uuid, body, user_id, thread_id, created_at) values (?, ?, ?, ?, ?)"
	stmtin, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmtin.Close()

	uuid := CreateUuID()
	stmtin.Exec(uuid, body, user.ID, conv.Id, time.Now())

	stmtout, err := DB.Prepare("select id, uuid, body, user_id, thread_id, created_at from posts where uuid = ?")
	if err != nil {
		return
	}
	defer stmtout.Close()

	err = stmtout.QueryRow(uuid).Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt)
	return
}
