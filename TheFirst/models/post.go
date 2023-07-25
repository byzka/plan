package models

import "time"

type Post struct {
	Id        int
	Uuid      string
	Body      string
	UserId    int
	ThreadId  int
	CreatedAt time.Time
}

func (post *Post) CreatedAtDate() string {
	return post.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// 找帖子主人
func (post *Post) User() (user User) {
	user = User{}
	DB.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = ?", post.UserId).Scan(
		&user.ID, &user.UUID, &user.Name, &user.Email, &user.CreatAT)
	return
}
