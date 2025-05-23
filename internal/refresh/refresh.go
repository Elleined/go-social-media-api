package refresh

import "time"

type Token struct {
	Id        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	Token     string    `json:"token" db:"token"`
	ExpiresAt time.Time `json:"expires_at" db:"expires_at"`
	RevokedAt time.Time `json:"revoked_at" db:"revoked_at"`
	UserId    int       `json:"user_id" db:"user_id"`
}
