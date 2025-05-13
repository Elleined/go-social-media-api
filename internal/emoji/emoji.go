package emoji

type Emoji struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
