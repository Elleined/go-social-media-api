package social

import "github.com/jmoiron/sqlx"

type (
	Repository interface {
		save(providerTypeId, providerId, userId int, signUpEmail string) (id int64, err error)

		findByUserId(userId int) (Social, error)
		findByProviderTypeAndId(providerTypeId, providerId int) (Social, error)
	}

	RepositoryImpl struct {
		*sqlx.DB
	}
)

func (r RepositoryImpl) findByUserId(userId int) (Social, error) {
	var social Social
	err := r.Get(&social, "SELECT * FROM social WHERE user_id = ?", userId)
	if err != nil {
		return social, err
	}

	return social, err
}

func (r RepositoryImpl) save(providerTypeId, providerId, userId int, signUpEmail string) (id int64, err error) {
	result, err := r.NamedExec("INSERT INTO user_social(provider_type_id, provider_id, user_id, sign_up_email) VALUES (:providerTypeId, :providerId, :userId, :signUpEmail)", map[string]any{
		"providerTypeId": providerTypeId,
		"providerId":     providerId,
		"userId":         userId,
		"signUpEmail":    signUpEmail,
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
