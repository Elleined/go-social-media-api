package emoji

import (
	"github.com/jmoiron/sqlx"
)

type (
	Repository interface {
		save(name string) (id int64, err error)

		findById(emojiId int) (Emoji, error)
		findByName(name string) (Emoji, error)
		findAll() ([]Emoji, error)

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

func (repository RepositoryImpl) findById(emojiId int) (Emoji, error) {
	var emoji Emoji
	err := repository.Get(&emoji, "SELECT * FROM emoji WHERE id = ?", emojiId)
	if err != nil {
		return Emoji{}, err
	}

	return emoji, nil
}

func (repository RepositoryImpl) findByName(name string) (Emoji, error) {
	var emoji Emoji
	err := repository.Get(&emoji, "SELECT * FROM emoji WHERE name = ?", name)
	if err != nil {
		return Emoji{}, err
	}

	return emoji, nil
}

func (repository RepositoryImpl) findAll() ([]Emoji, error) {
	emojis := make([]Emoji, 10)

	err := repository.Select(&emojis, "SELECT * FROM emoji")
	if err != nil {
		return nil, err
	}

	return emojis, nil
}

func (repository RepositoryImpl) isAlreadyExists(name string) (bool, error) {
	var exists bool
	err := repository.Get(&exists, "SELECT EXISTS(SELECT 1 FROM emoji WHERE name = ?)", name)
	if err != nil {
		return false, err
	}

	return exists, nil
}
