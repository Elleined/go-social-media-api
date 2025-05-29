package social_login

import "time"

type Social struct {
	Id             int       `json:"id" db:"id"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	ProviderId     string    `json:"provider_id" db:"provider_id"`
	UserId         int       `json:"user_id" db:"user_id"`
	ProviderTypeId int       `json:"provider_type_id" db:"provider_type_id"`
}
