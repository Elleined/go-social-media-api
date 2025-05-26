package refresh

import (
	"errors"
	"log"
	"strings"
)

type (
	Service interface {
		isValid(token Token) error

		Save(userId int) (token string, err error)

		getBy(token string) (Token, error)
		getAllBy(userId int) ([]Token, error)

		revoke(id int, userId int) (affectedRows int64, err error)
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

func (s ServiceImpl) isValid(token Token) error {
	if token.IsRevoked() {
		log.Print("Accessing illegal token") // add a panic, recover, and defer here to log
		return errors.New("token is revoked")
	}

	if token.IsExpired() {
		return errors.New("token is expired")
	}

	return nil
}

func (s ServiceImpl) Save(userId int) (token string, err error) {
	if userId <= 0 {
		return "", errors.New("userId is invalid")
	}

	token, err = s.repository.save(userId)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s ServiceImpl) getBy(token string) (Token, error) {
	if strings.TrimSpace(token) == "" {
		return Token{}, errors.New("token is empty")
	}

	refresh, err := s.repository.findBy(token)
	if err != nil {
		return Token{}, err
	}

	return refresh, nil
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

func (s ServiceImpl) revoke(id int, userId int) (affectedRows int64, err error) {
	if id <= 0 {
		return 0, errors.New("token is empty")
	}

	affectedRows, err = s.repository.revoke(id, userId)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("no affected rows")
	}

	return affectedRows, nil
}
