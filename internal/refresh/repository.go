package refresh

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type (
	Repository interface {
		Save(token string, userId int) (id int64, err error)
		FindByToken(token string, userId int) (Token, error)
		Delete(token string, userId int) (affectedRows int64, err error)
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

func (repository RepositoryImpl) Save(token string, userId int) (id int64, err error) {
	result, err := repository.NamedExec("INSERT INTO refresh_token(token, expires_at, user_id) VALUES (:token, :expiresAt, :userId)", map[string]any{
		"token":     token,
		"expiresAt": time.Now().AddDate(0, 1, 0),
		"userId":    userId,
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

func (repository RepositoryImpl) FindByToken(token string, userId int) (Token, error) {
	var result Token
	err := repository.Get(&result, "SELECT * FROM refresh_token WHERE token = ? AND user_id = ?", token, userId)
	if err != nil {
		return Token{}, err
	}

	return result, nil
}

func (repository RepositoryImpl) Delete(token string, userId int) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE refresh_token SET expires_at = NOW() WHERE token = :token AND user_id = :userId", map[string]any{
		"token":  token,
		"userId": userId,
	})
	if err != nil {
		return 0, err
	}

	affectedRows, err = result.RowsAffected()
	if err != nil {
		return 0, err
	}

	return affectedRows, err
}
