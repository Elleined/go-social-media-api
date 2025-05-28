package provider_type

type ProviderType struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
