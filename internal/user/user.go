package user

import "time"

type User struct {
	Id         int       `json:"id" db:"id"`
	CreatedAt  time.Time `json:"created_at" db:"created_at"`
	FirstName  string    `json:"first_name" db:"first_name"`
	LastName   string    `json:"last_name" db:"last_name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"-" db:"password"`
	Attachment string    `json:"attachment" db:"attachment"`
	IsActive   bool      `json:"is_active" db:"is_active"`
}
