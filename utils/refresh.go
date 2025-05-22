package utils

import (
	"github.com/jmoiron/sqlx"
	"time"
)

type (
	Token struct {
		Id        int       `json:"id" db:"id"`
		CreatedAt time.Time `json:"created_at" db:"created_at"`
		Token     string    `json:"token" db:"token"`
		ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
		RevokedAt time.Time `json:"revoked_at" db:"revoked_at"`
		UserId    int       `json:"user_id" db:"user_id"`
	}

	Service interface {
		Save(token string, userId int) (id int64, err error)
		FindByToken(token string) (Token, error)
	}

	ServiceImpl struct {
		db *sqlx.DB
	}
)

func NewService(db *sqlx.DB) Service {
	return &ServiceImpl{
		db: db,
	}
}

func (s ServiceImpl) Save(token string, userId int) (id int64, err error) {
	result, err := s.db.NamedExec("INSERT INTO refresh_token(token, user_id) VALUES (:token, :userId)", map[string]any{
		"token":  token,
		"userId": userId,
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

func (s ServiceImpl) FindByToken(token string) (Token, error) {
	var result Token
	err := s.db.Get(&result, "SELECT * FROM refresh_token WHERE token = ?", token)
	if err != nil {
		return Token{}, err
	}

	return result, nil
}
