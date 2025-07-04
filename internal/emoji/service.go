package emoji

import (
	"errors"
	"strings"
)

type (
	Service interface {
		save(name string) (id int64, err error)

		getById(emojiId int) (Emoji, error)
		getByName(name string) (Emoji, error)
		getAll() ([]Emoji, error)
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

func (s ServiceImpl) save(name string) (id int64, err error) {
	if strings.TrimSpace(name) == "" {
		return 0, errors.New("name is required")
	}

	exists, err := s.repository.isAlreadyExists(name)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("name already exists")
	}

	id, err = s.repository.save(name)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) getById(emojiId int) (Emoji, error) {
	if emojiId <= 0 {
		return Emoji{}, errors.New("emojiId is required")
	}

	emoji, err := s.repository.findById(emojiId)
	if err != nil {
		return Emoji{}, err
	}

	return emoji, nil
}

func (s ServiceImpl) getByName(name string) (Emoji, error) {
	if strings.TrimSpace(name) == "" {
		return Emoji{}, errors.New("name is required")
	}

	emoji, err := s.repository.findByName(name)
	if err != nil {
		return Emoji{}, err
	}

	return emoji, nil
}

func (s ServiceImpl) getAll() ([]Emoji, error) {
	emojis, err := s.repository.findAll()
	if err != nil {
		return nil, err
	}

	return emojis, nil
}
