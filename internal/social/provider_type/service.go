package provider_type

import (
	"errors"
	"strings"
)

type (
	Service interface {
		save(name string) (id int64, err error)

		getById(id int) (ProviderType, error)
		getAll() ([]ProviderType, error)

		update(id int, name string) (affectedRows int64, err error)
		delete(id int) (affectedRows int64, err error)
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

func (s ServiceImpl) getById(id int) (ProviderType, error) {
	if id <= 0 {
		return ProviderType{}, errors.New("id is required")
	}

	providerType, err := s.repository.findById(id)
	if err != nil {
		return ProviderType{}, err
	}

	return providerType, nil
}

func (s ServiceImpl) getAll() ([]ProviderType, error) {
	providerTypes, err := s.repository.findAll()
	if err != nil {
		return nil, err
	}

	return providerTypes, nil
}

func (s ServiceImpl) update(id int, name string) (affectedRows int64, err error) {
	if id <= 0 {
		return 0, errors.New("id is required")
	}

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

	affectedRows, err = s.repository.update(id, name)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("no rows updated")
	}

	return affectedRows, nil
}

func (s ServiceImpl) delete(id int) (affectedRows int64, err error) {
	if id <= 0 {
		return 0, errors.New("id is required")
	}

	affectedRows, err = s.repository.delete(id)
	if err != nil {
		return 0, err
	}

	if affectedRows <= 0 {
		return 0, errors.New("no rows updated")
	}

	return affectedRows, nil
}
