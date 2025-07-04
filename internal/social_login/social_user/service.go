package social_user

import (
	"errors"
	"strings"
)

type (
	Service interface {
		Save(providerTypeId, userId int, providerId string) (id int64, err error)
		GetByProviderTypeAndId(providerTypeId int, providerId string) (Social, error)
		IsAlreadyExists(providerTypeId int, providerId string) (bool, error)
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

func (s ServiceImpl) Save(providerTypeId, userId int, providerId string) (id int64, err error) {
	if providerTypeId <= 0 {
		return 0, errors.New("provider type id is required")
	}

	if strings.TrimSpace(providerId) == "" {
		return 0, errors.New("provider id is required")
	}

	if userId <= 0 {
		return 0, errors.New("user id is required")
	}

	id, err = s.repository.save(providerTypeId, userId, providerId)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) GetByProviderTypeAndId(providerTypeId int, providerId string) (Social, error) {
	if providerTypeId <= 0 {
		return Social{}, errors.New("provider type id is required")
	}

	if strings.TrimSpace(providerId) == "" {
		return Social{}, errors.New("provider id is required")
	}

	social, err := s.repository.findByProviderTypeAndId(providerTypeId, providerId)
	if err != nil {
		return Social{}, err
	}

	return social, nil
}

func (s ServiceImpl) IsAlreadyExists(providerTypeId int, providerId string) (bool, error) {
	if providerTypeId <= 0 {
		return false, errors.New("provider type id is required")
	}

	if strings.TrimSpace(providerId) == "" {
		return false, errors.New("provider id is required")
	}

	exists, err := s.repository.isAlreadyExists(providerTypeId, providerId)
	if err != nil {
		return false, err
	}

	return exists, nil
}
