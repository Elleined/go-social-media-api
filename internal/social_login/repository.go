package social_login

import "github.com/jmoiron/sqlx"

type (
	Repository interface {
		save(providerTypeId, userId int, providerId string) (id int64, err error)
		findByProviderTypeAndId(providerTypeId int, providerId string) (Social, error)
		isAlreadyExists(providerTypeId int, providerId string) (bool, error)
	}

	RepositoryImpl struct {
		*sqlx.DB
	}
)

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		DB: db,
	}
}

func (r RepositoryImpl) save(providerTypeId, userId int, providerId string) (id int64, err error) {
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

func (r RepositoryImpl) findByProviderTypeAndId(providerTypeId int, providerId string) (Social, error) {
	var social Social
	err := r.Get(&social, "SELECT * FROM user_social WHERE provider_type_id = ? AND provider_id = ?", providerTypeId, providerId)
	if err != nil {
		return social, err
	}

	return social, nil
}

func (r RepositoryImpl) isAlreadyExists(providerTypeId int, providerId string) (bool, error) {
	var exists bool
	err := r.Get(&exists, "SELECT EXISTS(SELECT 1 FROM user_social WHERE provider_type_id = ? AND provider_id = ?)", providerTypeId, providerId)
	if err != nil {
		return exists, err
	}

	return exists, nil
}
