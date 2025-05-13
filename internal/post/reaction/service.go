package reaction

import "errors"

type (
	Service interface {
		save(reactorId, postId, emojiId int) (id int64, err error)

		findAll(postId int) ([]Reaction, error)
		findAllByEmoji(postId int, emojiId int) ([]Reaction, error)

		update(reactorId, postId, newEmojiId int) (affectedRows int64, err error)

		delete(reactorId, postId int) (affectedRows int64, err error)
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

func (s ServiceImpl) save(reactorId, postId, emojiId int) (id int64, err error) {
	if reactorId <= 0 {
		return 0, errors.New("reactor id is required")
	}

	if postId <= 0 {
		return 0, errors.New("post id is required")
	}

	if emojiId <= 0 {
		return 0, errors.New("emoji id is required")
	}

	isAlreadyReacted, err := s.repository.isAlreadyReacted(reactorId, postId)
	if err != nil {
		return 0, err
	}

	if isAlreadyReacted {
		return 0, errors.New("reactor already reacted")
	}

	id, err = s.repository.save(reactorId, postId, emojiId)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) findAll(postId int) ([]Reaction, error) {
	if postId <= 0 {
		return nil, errors.New("post id is required")
	}

	reactions, err := s.repository.findAll(postId)
	if err != nil {
		return nil, err
	}

	return reactions, nil
}

func (s ServiceImpl) findAllByEmoji(postId int, emojiId int) ([]Reaction, error) {
	if postId <= 0 {
		return nil, errors.New("post id is required")
	}

	if emojiId <= 0 {
		return nil, errors.New("emoji id is required")
	}

	reactions, err := s.repository.findAllByEmoji(postId, emojiId)
	if err != nil {
		return nil, err
	}

	return reactions, nil
}

func (s ServiceImpl) update(reactorId, postId, newEmojiId int) (affectedRows int64, err error) {
	if reactorId <= 0 {
		return 0, errors.New("reactor id is required")
	}

	if postId <= 0 {
		return 0, errors.New("post id is required")
	}

	if newEmojiId <= 0 {
		return 0, errors.New("new emoji id is required")
	}

	isAlreadyReacted, err := s.repository.isAlreadyReacted(reactorId, postId)
	if err != nil {
		return 0, err
	}

	if !isAlreadyReacted {
		return 0, errors.New("current user doesn't reacted to this post")
	}

	affectedRows, err = s.repository.update(reactorId, postId, newEmojiId)
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (s ServiceImpl) delete(reactorId, postId int) (affectedRows int64, err error) {
	if reactorId <= 0 {
		return 0, errors.New("reactor id is required")
	}

	if postId <= 0 {
		return 0, errors.New("post id is required")
	}

	isAlreadyReacted, err := s.repository.isAlreadyReacted(reactorId, postId)
	if err != nil {
		return 0, err
	}

	if !isAlreadyReacted {
		return 0, errors.New("current user doesn't reacted to this post")
	}

	affectedRows, err = s.repository.delete(reactorId, postId)
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}
