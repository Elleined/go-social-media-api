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
	//TODO implement me
	panic("implement me")
}

func (r RepositoryImpl) findAll() ([]Emoji, error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryImpl) update(emojiId int, newName string) (affectedRows int64, err error) {
	//TODO implement me
	panic("implement me")
}

func (r RepositoryImpl) delete(emojiId int) (affectedRows int64, err error) {
	//TODO implement me
	panic("implement me")
}
