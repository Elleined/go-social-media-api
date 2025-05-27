package refresh

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"os"
	"strconv"
	"time"
)

type (
	Repository interface {
		save(userId int) (token string, err error)
		findBy(token string) (Token, error)

		findAllBy(userId int) ([]Token, error)

		revoke(id int, userId int) (affectedRows int64, err error)
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

func (repository RepositoryImpl) save(userId int) (token string, err error) {
	tokenExpiration, err := strconv.Atoi(os.Getenv("REFRESH_TOKEN_EXPIRATION_IN_DAYS"))
	if err != nil {
		return "", err
	}

	token = uuid.New().String()
	_, err = repository.NamedExec("INSERT INTO refresh_token(token, expires_at, user_id) VALUES (:token, :expiresAt, :userId)", map[string]any{
		"token":     token,
		"expiresAt": time.Now().AddDate(0, 0, tokenExpiration),
		"userId":    userId,
	})
	if err != nil {
		return "", err
	}

	return token, nil
}

func (repository RepositoryImpl) findBy(token string) (Token, error) {
	var result Token
	err := repository.Get(&result, "SELECT * FROM refresh_token WHERE token = ?", token)
	if err != nil {
		return Token{}, err
	}

	return result, nil
}

func (repository RepositoryImpl) findAllBy(userId int) ([]Token, error) {
	tokens := make([]Token, 10)
	err := repository.Select(&tokens, "SELECT * FROM refresh_token WHERE user_id = ? ORDER BY created_at DESC", userId)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}

func (repository RepositoryImpl) revoke(id int, userId int) (affectedRows int64, err error) {
	result, err := repository.NamedExec("UPDATE refresh_token SET revoked_at = NOW() WHERE id = :id AND user_id = :userId", map[string]any{
		"id":     id,
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
