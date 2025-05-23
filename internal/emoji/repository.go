package emoji

import (
	"github.com/jmoiron/sqlx"
)

type (
	Repository interface {
		save(name string) (id int64, err error)
		findAll() ([]Emoji, error)
		update(emojiId int, newName string) (affectedRows int64, err error)
		delete(emojiId int) (affectedRows int64, err error)

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

func (repository RepositoryImpl) save(name string) (id int64, err error) {
	result, err := repository.NamedExec("INSERT INTO emoji (name) VALUES (:name)", map[string]any{
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

func (repository RepositoryImpl) findAll() ([]Emoji, error) {
	emojis := make([]Emoji, 10)

	err := repository.Select(&emojis, "SELECT * FROM emoji")
	if err != nil {
		return nil, err
	}

	return emojis, nil
}

func (repository RepositoryImpl) update(emojiId int, newName string) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE emoji SET name = :name WHERE id = :id", map[string]any{
		"name": newName,
		"id":   emojiId,
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

func (repository RepositoryImpl) delete(emojiId int) (affectedRows int64, err error) {
	result, err := repository.NamedExec("DELETE FROM emoji WHERE id = :id", map[string]any{
		"id": emojiId,
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

func (repository RepositoryImpl) isAlreadyExists(name string) (bool, error) {
	var exists bool
	err := repository.Get(&exists, "SELECT EXISTS(SELECT 1 FROM emoji WHERE name = ?)", name)
	if err != nil {
		return false, err
	}

	return exists, nil
}
