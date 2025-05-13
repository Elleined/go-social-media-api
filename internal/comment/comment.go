package comment

import (
	"database/sql"
	"time"
)

type Comment struct {
	Id         int            `json:"id" db:"id"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	Content    string         `json:"content" db:"content"`
	Attachment sql.NullString `json:"attachment"  db:"attachment"`
	IsDeleted  bool           `json:"is_deleted"  db:"is_deleted"`
	AuthorId   int            `json:"author_id"  db:"author_id"`
	PostId     int            `json:"post_id" db:"post_id"`
}
