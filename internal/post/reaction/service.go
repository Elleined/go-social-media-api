package reaction

import "errors"

type Service interface {
	save(reactorId, postId, emojiId int) (id int64, err error)

	findAll(postId int) ([]Reaction, error)
	findAllByEmoji(postId int, emojiId int) ([]Reaction, error)

	delete(reactorId, postId int) (affectedRows int64, err error)
}

type ServiceImpl struct {
	repository Repository
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

func (s ServiceImpl) delete(reactorId, postId int) (affectedRows int64, err error) {
	if reactorId <= 0 {
		return 0, errors.New("reactor id is required")
	}

	if postId <= 0 {
		return 0, errors.New("post id is required")
	}

	affectedRows, err = s.repository.delete(reactorId, postId)
	if err != nil {
		return 0, err
	}
	
	return affectedRows, nil
}

func NewService(repository Repository) Service {
	return &ServiceImpl{
		repository: repository,
	}
}
