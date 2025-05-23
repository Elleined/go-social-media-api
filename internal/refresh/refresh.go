package refresh

import (
	"database/sql"
	"time"
)

type Token struct {
	Id        int          `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	Token     string       `json:"token" db:"token"`
	ExpiresAt sql.NullTime `json:"expires_at" db:"expires_at"`
	RevokedAt sql.NullTime `json:"revoked_at" db:"revoked_at"`
	UserId    int          `json:"user_id" db:"user_id"`
}

func (t Token) IsExpired() bool {
	return !t.ExpiresAt.Valid || time.Now().After(t.ExpiresAt.Time)
}

func (t Token) IsRevoked() bool {
	return t.RevokedAt.Valid // If theres a value it is revoked
}
