package emoji

import (
	"errors"
	"strings"
)

type (
	Service interface {
		save(name string) (id int64, err error)
		getById(emojiId int) (Emoji, error)
		getAll() ([]Emoji, error)
		update(emojiId int, newName string) (affectedRows int64, err error)
		delete(emojiId int) (affectedRows int64, err error)
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

func (s ServiceImpl) getAll() ([]Emoji, error) {
	emojis, err := s.repository.findAll()
	if err != nil {
		return nil, err
	}

	return emojis, nil
}

func (s ServiceImpl) update(emojiId int, newName string) (affectedRows int64, err error) {
	if emojiId <= 0 {
		return 0, errors.New("emojiId is required")
	}

	if strings.TrimSpace(newName) == "" {
		return 0, errors.New("name is required")
	}

	exists, err := s.repository.isAlreadyExists(newName)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("name already exists")
	}

	affectedRows, err = s.repository.update(emojiId, newName)
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (s ServiceImpl) delete(emojiId int) (affectedRows int64, err error) {
	if emojiId <= 0 {
		return 0, errors.New("emojiId is required")
	}

	affectedRows, err = s.repository.delete(emojiId)
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}
