package social

import "github.com/jmoiron/sqlx"

type (
	Repository interface {
		save(providerTypeId, providerId, userId int) (id int64, err error)
		findByProviderTypeAndId(providerTypeId, providerId int) (Social, error)
		isAlreadyExists(providerTypeId, providerId int) (bool, error)
	}

	RepositoryImpl struct {
		*sqlx.DB
	}
)

func (r RepositoryImpl) isAlreadyExists(providerTypeId, providerId int) (bool, error) {
	var exists bool
	err := r.Get(&exists, "SELECT EXISTS(SELECT 1 FROM user_social WHERE provider_type_id = ? AND provider_id = ?)", providerTypeId, providerId)
	if err != nil {
		return exists, err
	}

	return exists, nil
}

func (r RepositoryImpl) save(providerTypeId, providerId, userId int) (id int64, err error) {
	result, err := r.NamedExec("INSERT INTO user_social(provider_type_id, provider_id, user_id) VALUES (:providerTypeId, :providerId, :userId)", map[string]any{
		"providerTypeId": providerTypeId,
		"providerId":     providerId,
		"userId":         userId,
	})
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r RepositoryImpl) findByProviderTypeAndId(providerTypeId, providerId int) (Social, error) {
	var social Social
	err := r.Get(&social, "SELECT * FROM user_social WHERE provider_type_id = ? AND provider_id = ?", providerTypeId, providerId)
	if err != nil {
		return social, err
	}

	return social, nil
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		DB: db,
	}
}
