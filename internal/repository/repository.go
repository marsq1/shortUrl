package repository

import "github.com/jmoiron/sqlx"

type ShorterUrl interface {
	SaveShortUrl (hashUrl, originalUrl string)
	GetOriginalUrl(shortUrl string) string
}

type Repository struct {
	ShorterUrl
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{ShorterUrl: NewUrlRepos(db)}
}