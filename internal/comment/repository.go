package comment

import "github.com/jmoiron/sqlx"

type Repository interface {
}

type RepositoryImpl struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) Repository {
	return &RepositoryImpl{
		db: db,
	}
}
