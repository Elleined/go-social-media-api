package commentreaction

import (
	"errors"
	"social-media-application/internal/paging"
)

type (
	Service interface {
		save(reactorId, postId, commentId, emojiId int) (id int64, err error)

		getById(postId, commentId, reactionId int) (Reaction, error)
		getAll(postId, commentId int, request *paging.PageRequest) (*paging.Page[Reaction], error)
		getAllByEmoji(postId, commentId, emojiId int, request *paging.PageRequest) (*paging.Page[Reaction], error)

		update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error)

		delete(reactorId, postId, commentId int) (affectedRows int64, err error)
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
	if reactorId <= 0 {
		return 0, errors.New("reactor is required")
	}

	if commentId <= 0 {
		return 0, errors.New("commentId is required")
	}

	if emojiId <= 0 {
		return 0, errors.New("emojiId is required")
	}

	isAlreadyReacted, err := s.repository.isAlreadyReacted(reactorId, postId, commentId)
	if err != nil {
		return 0, err
	}

	if isAlreadyReacted {
		return 0, errors.New("reactor already reacted")
	}

	id, err = s.repository.save(reactorId, postId, commentId, emojiId)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) getById(postId, commentId, reactionId int) (Reaction, error) {
	if postId <= 0 {
		return Reaction{}, errors.New("postId is required")
	}

	if commentId <= 0 {
		return Reaction{}, errors.New("commentId is required")
	}

	if reactionId <= 0 {
		return Reaction{}, errors.New("reactionId is required")
	}

	reaction, err := s.repository.findById(postId, commentId, reactionId)
	if err != nil {
		return Reaction{}, err
	}

	return reaction, nil
}

func (s ServiceImpl) getAll(postId, commentId int, request *paging.PageRequest) (*paging.Page[Reaction], error) {
	if postId <= 0 {
		return nil, errors.New("postId is required")
	}

	if commentId <= 0 {
		return nil, errors.New("commentId is required")
	}

	reactions, err := s.repository.findAll(postId, commentId, request)
	if err != nil {
		return nil, err
	}

	return reactions, nil
}

func (s ServiceImpl) getAllByEmoji(postId, commentId, emojiId int, request *paging.PageRequest) (*paging.Page[Reaction], error) {
	if postId <= 0 {
		return nil, errors.New("postId is required")
	}

	if commentId <= 0 {
		return nil, errors.New("commentId is required")
	}

	if emojiId <= 0 {
		return nil, errors.New("emojiId is required")
	}

	reactions, err := s.repository.findAllByEmoji(postId, commentId, emojiId, request)
	if err != nil {
		return nil, err
	}

	return reactions, nil
}

func (s ServiceImpl) update(reactorId, postId, commentId, newEmojiId int) (affectedRows int64, err error) {
	if reactorId <= 0 {
		return 0, errors.New("reactor is required")
	}

	if postId <= 0 {
		return 0, errors.New("postId is required")
	}

	if commentId <= 0 {
		return 0, errors.New("commentId is required")
	}

	if newEmojiId <= 0 {
		return 0, errors.New("emojiId is required")
	}

	isAlreadyReacted, err := s.repository.isAlreadyReacted(reactorId, postId, commentId)
	if err != nil {
		return 0, err
	}

	if !isAlreadyReacted {
		return 0, errors.New("reactor does not reacted")
	}

	affectedRows, err = s.repository.update(reactorId, postId, commentId, newEmojiId)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("no affected rows")
	}

	return affectedRows, nil
}

func (s ServiceImpl) delete(reactorId, postId, commentId int) (affectedRows int64, err error) {
	if reactorId <= 0 {
		return 0, errors.New("reactor is required")
	}

	if postId <= 0 {
		return 0, errors.New("postId is required")
	}

	if commentId <= 0 {
		return 0, errors.New("commentId is required")
	}

	isAlreadyReacted, err := s.repository.isAlreadyReacted(reactorId, postId, commentId)
	if err != nil {
		return 0, err
	}

	if !isAlreadyReacted {
		return 0, errors.New("reactor does not reacted")
	}

	affectedRows, err = s.repository.delete(reactorId, postId, commentId)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("no affected rows")
	}

	return affectedRows, nil
}
