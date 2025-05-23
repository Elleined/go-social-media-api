package refresh

import (
	"errors"
	"log"
	"strings"
)

type (
	Service interface {
		isValid(token string, userId int) (bool, error)
		save(token string, userId int) (id int64, err error)

		getAllBy(userId int) ([]Token, error)

		revoke(token string, userId int) (affectedRows int64, err error)
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

func (s ServiceImpl) isValid(token string, userId int) (bool, error) {
	if strings.TrimSpace(token) == "" {
		return false, errors.New("token is empty")
	}

	if userId <= 0 {
		return false, errors.New("userId is invalid")
	}

	refresh, err := s.repository.findBy(token, userId)
	if err != nil {
		return false, err
	}

	if refresh.IsRevoked() {
		log.Print("Accessing illegal token") // add a panic, recover, and defer here to log
		return false, errors.New("token is revoked")
	}

	if refresh.IsExpired() {
		return false, errors.New("token is expired")
	}

	return true, nil
}

func (s ServiceImpl) save(token string, userId int) (id int64, err error) {
	if strings.TrimSpace(token) == "" {
		return 0, errors.New("token is empty")
	}

	if userId <= 0 {
		return 0, errors.New("userId is invalid")
	}

	id, err = s.save(token, userId)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) getAllBy(userId int) ([]Token, error) {
	if userId <= 0 {
		return nil, errors.New("userId is invalid")
	}

	refreshTokens, err := s.repository.findAllBy(userId)
	if err != nil {
		return nil, err
	}

	return refreshTokens, nil
}

func (s ServiceImpl) revoke(token string, userId int) (affectedRows int64, err error) {
	if strings.TrimSpace(token) == "" {
		return 0, errors.New("token is empty")
	}

	if userId <= 0 {
		return 0, errors.New("userId is invalid")
	}

	affectedRows, err = s.repository.delete(token, userId)
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}
