package social

import "github.com/jmoiron/sqlx"

type (
	Repository interface {
		save(providerTypeId, providerId, userId int, emailAtSignup string) (id int64, err error)
		findByProviderTypeAndId(providerTypeId, providerId int) (Social, error)
	}

	RepositoryImpl struct {
		*sqlx.DB
	}
)

func (r RepositoryImpl) save(providerTypeId, providerId, userId int, emailAtSignup string) (id int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryImpl) findByProviderTypeAndId(providerTypeId, providerId int) (Social, error) {
	//TODO implement me
	panic("implement me")
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		DB: db,
	}
}
