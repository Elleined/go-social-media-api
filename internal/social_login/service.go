package social_login

import "errors"

type (
	Service interface {
		Save(providerTypeId, providerId, userId int) (id int64, err error)
		GetByProviderTypeAndId(providerTypeId, providerId int) (Social, error)
		IsAlreadyExists(providerTypeId, providerId int) (bool, error)
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

func (s ServiceImpl) Save(providerTypeId, providerId, userId int) (id int64, err error) {
	if providerTypeId <= 0 {
		return 0, errors.New("provider type id is required")
	}

	if providerId <= 0 {
		return 0, errors.New("provider id is required")
	}

	if userId <= 0 {
		return 0, errors.New("user id is required")
	}

	id, err = s.repository.save(providerTypeId, providerId, userId)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s ServiceImpl) GetByProviderTypeAndId(providerTypeId, providerId int) (Social, error) {
	if providerTypeId <= 0 {
		return Social{}, errors.New("provider type id is required")
	}

	if providerId <= 0 {
		return Social{}, errors.New("provider id is required")
	}

	social, err := s.repository.findByProviderTypeAndId(providerTypeId, providerId)
	if err != nil {
		return Social{}, err
	}

	return social, nil
}

func (s ServiceImpl) IsAlreadyExists(providerTypeId, providerId int) (bool, error) {
	if providerTypeId <= 0 {
		return false, errors.New("provider type id is required")
	}

	if providerId <= 0 {
		return false, errors.New("provider id is required")
	}

	exists, err := s.repository.isAlreadyExists(providerTypeId, providerId)
	if err != nil {
		return false, err
	}

	return exists, nil
}
