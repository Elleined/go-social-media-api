package commentreaction

import (
	"github.com/jmoiron/sqlx"
)

type (
	Repository interface {
		save(reactorId, postId, commentId, emojiId int) (id int64, err error)

		findAll(postId, commentId int) ([]Reaction, error)
		findAllByEmoji(postId, commentId, emojiId int) ([]Reaction, error)

		update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error)

		delete(reactorId, postId, commentId int) (affectedRows int64, err error)

		isAlreadyReacted(reactorId, postId, commentId int) (bool, error)
	}

	RepositoryImpl struct {
		db *sqlx.DB
	}
)

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r RepositoryImpl) save(reactorId, postId, commentId, emojiId int) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryImpl) findAll(postId, commentId int) ([]Reaction, error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryImpl) findAllByEmoji(postId, commentId, emojiId int) ([]Reaction, error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryImpl) update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryImpl) delete(reactorId, postId, commentId int) (affectedRows int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryImpl) isAlreadyReacted(reactorId, postId, commentId int) (bool, error) {
	//TODO implement me
	panic("implement me")
}
