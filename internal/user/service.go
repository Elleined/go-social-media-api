package user

import (
	"errors"
	"social-media-application/internal/paging"
	pd "social-media-application/internal/user/password"
	"strings"
)

type (
	Service interface {
		saveLocal(firstName, lastName, email, password, attachment string) (id int64, err error)
		SaveSocial(firstName, lastName, email string) (id int64, err error) // for social register

		getById(id int) (User, error)
		GetByEmail(email string) (User, error)

		getAll(isActive bool, request *paging.PageRequest) (*paging.Page[User], error)

		deleteById(id int) (affectedRows int64, err error)

		changeAttachment(userId int, attachment string) (affectedRows int64, err error)
		changeStatus(userId int, isActive bool) (affectedRows int64, err error)
		changePassword(userId int, newPassword string) (affectedRows int64, err error)
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

func (s ServiceImpl) saveLocal(firstName, lastName, email, password, attachment string) (id int64, err error) {
	if strings.TrimSpace(firstName) == "" {
		return 0, errors.New("first name is required")
	}

	if strings.TrimSpace(lastName) == "" {
		return 0, errors.New("last name is required")
	}

	if strings.TrimSpace(email) == "" {
		return 0, errors.New("email is required")
	}

	if strings.TrimSpace(password) == "" {
		return 0, errors.New("password is required")
	}

	exists, err := s.repository.isEmailExists(email)
	if err != nil {
		return 0, err
	}

	if exists {
		return 0, errors.New("email already exists")
	}

	hashedPassword, err := pd.Encrypt(password)
	if err != nil {
		return 0, err
	}

	id, err = s.repository.saveLocal(firstName, lastName, email, hashedPassword, attachment)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) SaveSocial(firstName, lastName, email string) (id int64, err error) {
	if strings.TrimSpace(firstName) == "" {
		return 0, errors.New("first name is required")
	}

	if strings.TrimSpace(lastName) == "" {
		return 0, errors.New("last name is required")
	}

	if strings.TrimSpace(email) == "" {
		return 0, errors.New("email is required")
	}

	id, err = s.repository.saveSocial(firstName, lastName, email)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) getById(id int) (User, error) {
	if id <= 0 {
		return User{}, errors.New("user id is required")
	}

	user, err := s.repository.findById(id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s ServiceImpl) GetByEmail(email string) (User, error) {
	if strings.TrimSpace(email) == "" {
		return User{}, errors.New("email is required")
	}

	user, err := s.repository.findByEmail(email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s ServiceImpl) getAll(isActive bool, request *paging.PageRequest) (*paging.Page[User], error) {
	users, err := s.repository.findAll(isActive, request)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s ServiceImpl) deleteById(id int) (affectedRows int64, err error) {
	if id <= 0 {
		return 0, errors.New("user id is required")
	}

	affectedRows, err = s.repository.deleteById(id)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("no rows affected")
	}

	return affectedRows, nil
}

func (s ServiceImpl) changeAttachment(userId int, attachment string) (affectedRows int64, err error) {
	if userId <= 0 {
		return 0, errors.New("user id is required")
	}

	if strings.TrimSpace(attachment) == "" {
		return 0, errors.New("attachment is required")
	}

	affectedRows, err = s.repository.changeAttachment(userId, attachment)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("no rows affected")
	}

	return affectedRows, nil
}

func (s ServiceImpl) changeStatus(userId int, isActive bool) (affectedRows int64, err error) {
	if userId <= 0 {
		return 0, errors.New("user id is required")
	}

	affectedRows, err = s.repository.changeStatus(userId, isActive)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("no rows affected")
	}

	return affectedRows, nil
}

func (s ServiceImpl) changePassword(userId int, newPassword string) (affectedRows int64, err error) {
	if userId <= 0 {
		return 0, errors.New("user id is required")
	}

	if strings.TrimSpace(newPassword) == "" {
		return 0, errors.New("new password is required")
	}

	hashedPassword, err := pd.Encrypt(newPassword)
	if err != nil {
		return 0, err
	}

	affectedRows, err = s.repository.changePassword(userId, hashedPassword)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("no rows affected")
	}

	return affectedRows, nil
}
