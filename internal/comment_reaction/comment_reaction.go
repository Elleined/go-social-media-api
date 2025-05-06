package commentreaction

import (
	"time"
)

type CommentReaction struct {
	Id        int       `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	ReactorId int       `json:"reactor_id"`
	CommentId int       `json:"comment_id"`
	EmojiId   int       `json:"emoji_id"`
}
