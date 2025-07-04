package provider_type

import "github.com/jmoiron/sqlx"

type (
	Repository interface {
		save(name string) (id int64, err error)

		findById(id int) (ProviderType, error)
		findByName(name string) (ProviderType, error)
		findAll() ([]ProviderType, error)

		isAlreadyExists(name string) (bool, error)
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

func (r RepositoryImpl) save(name string) (id int64, err error) {
	result, err := r.NamedExec("INSERT INTO provider_type(name) VALUES (:name)", map[string]any{
		"name": name,
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

func (r RepositoryImpl) findById(id int) (ProviderType, error) {
	var providerType ProviderType
	err := r.Get(&providerType, `SELECT * FROM provider_type WHERE id = ?`, id)
	if err != nil {
		return ProviderType{}, err
	}

	return providerType, nil
}

func (r RepositoryImpl) findByName(name string) (ProviderType, error) {
	var providerType ProviderType
	err := r.Get(&providerType, `SELECT * FROM provider_type WHERE name = ?`, name)
	if err != nil {
		return ProviderType{}, err
	}

	return providerType, nil
}

func (r RepositoryImpl) findAll() ([]ProviderType, error) {
	providerTypes := make([]ProviderType, 10)
	err := r.Select(&providerTypes, `SELECT * FROM provider_type ORDER BY name`)
	if err != nil {
		return nil, err
	}

	return providerTypes, nil
}

func (r RepositoryImpl) isAlreadyExists(name string) (bool, error) {
	var exists bool
	err := r.Get(&exists, `SELECT EXISTS(SELECT 1 FROM provider_type WHERE name = ?)`, name)
	if err != nil {
		return exists, err
	}

	return exists, err
}
