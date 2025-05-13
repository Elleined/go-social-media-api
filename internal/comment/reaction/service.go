package commentreaction

type (
	Service interface {
		save(reactorId, postId, commentId, emojiId int) (id int64, err error)

		getAll(postId, commentId int) ([]Reaction, error)
		getAllByEmoji(postId, commentId, emojiId int) ([]Reaction, error)

		update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error)

		delete(reactorId, postId, commentId int) (affectedRows int64, err error)

		isAlreadyReacted(reactorId, postId, commentId int) (bool, error)
	}

	ServiceImpl struct {
		repository Repository
	}
)

func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}

func (s ServiceImpl) save(reactorId, postId, commentId, emojiId int) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) getAll(postId, commentId int) ([]Reaction, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) getAllByEmoji(postId, commentId, emojiId int) ([]Reaction, error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) delete(reactorId, postId, commentId int) (affectedRows int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (s ServiceImpl) isAlreadyReacted(reactorId, postId, commentId int) (bool, error) {
	//TODO implement me
	panic("implement me")
}
