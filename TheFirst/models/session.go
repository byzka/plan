package models

import "time"

type Session struct {
	Id        int
	Uuid      string
	Age       int
	Email     string
	UserId    int
	CreatedAt time.Time
}

// 检测数据库中的Session是否有效
func (session *Session) Check_session() (vaild bool, err error) {
	err = DB.QueryRow("SELECT id, uuid, email, user_id, created_at FROM sessions WHERE uuid = ?", session.Uuid).
		Scan(&session.Id, &session.Uuid, &session.Email, &session.UserId, &session.CreatedAt)
	if err != nil {
		vaild = false
		return
	}
	if session.Id != 0 {
		vaild = true
	}
	return
}

// 从数据库中删掉对应的session
func (session *Session) DeleteByUUID() (err error) {
	statement := "delete from sessions where uuid = ?"
	stmt, err := DB.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(session.Uuid)
	return
}

// 从数据库中删掉所有session
func (session *Session) Deleteall() (err error) {
	statement := "delete from sessions"
	_, err = DB.Exec(statement)
	return
}

// 从session 中获取用户
func (session *Session) User() (user User, err error) {
	user = User{}
	err = DB.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", session.UserId).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatAT)
	return
}
