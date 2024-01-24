package store

type ShortURL struct {
	ID          int    `json:"id"`
	Destination string `json:"destination"`
	Slug        string `json:"slug"`
}
