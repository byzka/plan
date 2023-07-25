package models

import "time"

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

// 设置 CreatedAt 日期的格式
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// 获取线程中帖子数量
func (thread *Thread) NumReplies() (count int) {
	rows, err := DB.Query("SELECT count(*) FROM posts where thread_id = ?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// 获取帖子
func (thread *Thread) Posts() (posts []Post, err error) {
	rows, err := DB.Query("SELECT id, uuid, body, user_id, thread_id, created_at FROM posts where thread_id = ?", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		if err = rows.Scan(&post.Id, &post.Uuid, &post.Body, &post.UserId, &post.ThreadId, &post.CreatedAt); err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// 获取数据库中所有线程并返回
func Threads() (threads []Thread, err error) {
	rows, err := DB.Query("SELECT id, uuid, topic, user_id, created_at FROM threads ORDER BY created_at DESC")
	if err != nil {
		return
	}
	for rows.Next() {
		conv := Thread{}
		if err = rows.Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt); err != nil {
			return
		}
		threads = append(threads, conv)
	}
	rows.Close()
	return
}

// 通过UUID查找对应线程
func ThreadByUUID(uuid string) (conv Thread, err error) {
	conv = Thread{}
	err = DB.QueryRow("SELECT id, uuid, topic, user_id, created_at FROM threads WHERE uuid = ?", uuid).
		Scan(&conv.Id, &conv.Uuid, &conv.Topic, &conv.UserId, &conv.CreatedAt)
	return
}

// 找到谁开的线程
func (thread *Thread) User() (user User) {
	user = User{}
	DB.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", thread.UserId).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatAT)
	return
}
