package post

import (
	"database/sql"
	"time"
)

type Post struct {
	Id         int            `json:"id" db:"id"`
	CreatedAt  time.Time      `json:"created_at" db:"created_at"`
	Subject    string         `json:"subject" db:"subject"`
	Content    string         `json:"content" db:"content"`
	Attachment sql.NullString `json:"attachment" db:"attachment"`
	IsDeleted  bool           `json:"-" db:"is_deleted"`
	AuthorId   int            `json:"author_id" db:"author_id"`
}
