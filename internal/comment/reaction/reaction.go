package commentreaction

import (
	"time"
)

type Reaction struct {
	Id        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	ReactorId int       `json:"reactor_id" db:"reactor_id"`
	CommentId int       `json:"comment_id" db:"comment_id"`
	EmojiId   int       `json:"emoji_id"  db:"emoji_id"`
}
