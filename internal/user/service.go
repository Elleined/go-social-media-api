package user

import (
	"errors"
	pd "social-media-application/internal/user/password"
	"social-media-application/utils"
	"strings"
)

type (
	Service interface {
		save(firstName, lastName, email, password string) (id int64, err error)

		getById(id int) (User, error)
		getByEmail(email string) (User, error)

		getAll(isActive bool, limit, offset int) ([]User, error)

		deleteById(id int) (affectedRows int64, err error)

		changeStatus(userId int, isActive bool) (affectedRows int64, err error)
		changePassword(userId int, newPassword string) (affectedRows int64, err error)

		login(username, password string) (jwt string, err error)
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

func (s ServiceImpl) save(firstName, lastName, email, password string) (id int64, err error) {
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

	id, err = s.repository.save(firstName, lastName, email, hashedPassword)
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

func (s ServiceImpl) getByEmail(email string) (User, error) {
	if strings.TrimSpace(email) == "" {
		return User{}, errors.New("email is required")
	}

	user, err := s.repository.findByEmail(email)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (s ServiceImpl) getAll(isActive bool, limit, offset int) ([]User, error) {
	if limit < 0 {
		return nil, errors.New("limit is required")
	}

	if offset < 0 {
		return nil, errors.New("offset is required")
	}

	users, err := s.repository.findAll(isActive, limit, offset)
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

	return affectedRows, nil
}

func (s ServiceImpl) login(username string, password string) (jwt string, err error) {
	if strings.TrimSpace(username) == "" {
		return "", errors.New("username is required")
	}

	if strings.TrimSpace(password) == "" {
		return "", errors.New("password is required")
	}

	user, err := s.repository.findByEmail(username)
	if err != nil {
		return "", errors.New("invalid credentials " + err.Error())
	}

	if pd.IsPasswordMatch(password, user.Password) {
		return "", errors.New("invalid credentials ")
	}

	user, err = s.getByEmail(user.Email)
	if err != nil {
		return "", errors.New("invalid credentials" + err.Error())
	}

	jwt, err = utils.GenerateJWT(user.Id)
	if err != nil {
		return "", errors.New("cannot generate jwt: " + err.Error())
	}

	return jwt, nil
}
