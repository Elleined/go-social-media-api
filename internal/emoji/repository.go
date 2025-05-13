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
	}

	RepositoryImpl struct {
		db *sqlx.DB
	}
)

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}

func (r RepositoryImpl) save(name string) (id int64, err error) {
	result, err := r.db.NamedExec("INSERT INTO emoji (name) VALUES (:name)", map[string]any{
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

func (r RepositoryImpl) findAll() ([]Emoji, error) {
	emojis := make([]Emoji, 10)

	err := r.db.Select(&emojis, "SELECT * FROM emoji")
	if err != nil {
		return nil, err
	}

	return emojis, nil
}

func (r RepositoryImpl) update(emojiId int, newName string) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("UPDATE emoji SET name = :name WHERE id = :id", map[string]any{
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

func (r RepositoryImpl) delete(emojiId int) (affectedRows int64, err error) {
	result, err := r.db.NamedExec("DELETE FROM emoji WHERE id = :id", map[string]any{
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
