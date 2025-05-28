package social

import "github.com/jmoiron/sqlx"

type (
	Repository interface {
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
