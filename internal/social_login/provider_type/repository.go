package provider_type

import "github.com/jmoiron/sqlx"

type (
	Repository interface {
		save(name string) (id int64, err error)

		findById(id int) (ProviderType, error)
		findByName(name string) (ProviderType, error)

		findAll() ([]ProviderType, error)

		update(id int, name string) (affectedRows int64, err error)
		delete(id int) (affectedRows int64, err error)

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

func (r RepositoryImpl) update(id int, name string) (affectedRows int64, err error) {
	result, err := r.NamedExec("UPDATE provider_type SET name = :name WHERE id = :id", map[string]any{
		"id":   id,
		"name": name,
	})
	if err != nil {
		return 0, err
	}

	affectedRows, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (r RepositoryImpl) delete(id int) (affectedRows int64, err error) {
	result, err := r.NamedExec("DELETE FROM provider_type WHERE id = :id", map[string]any{
		"id": id,
	})
	if err != nil {
		return 0, err
	}

	affectedRows, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, nil
}

func (r RepositoryImpl) isAlreadyExists(name string) (bool, error) {
	var exists bool
	err := r.Get(&exists, `SELECT EXISTS(SELECT 1 FROM provider_type WHERE name = ?)`, name)
	if err != nil {
		return exists, err
	}

	return exists, err
}
