package comment

import (
	"time"
)

type Comment struct {
	Id         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	Content    string    `json:"content"`
	Attachment string    `json:"attachment"`
	IsDeleted  bool      `json:"is_deleted"`
	AuthorId   int       `json:"author_id"`
	PostId     int       `json:"post_id"`
}
