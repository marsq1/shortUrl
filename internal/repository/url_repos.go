package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const (
	tableName = "save"
	port = "http://localhost:8000/"
)

type UrlRepos struct {
	db *sqlx.DB
}

func NewUrlRepos(db *sqlx.DB) *UrlRepos {
	return &UrlRepos{db: db}
}

func (r *UrlRepos) SaveShortUrl (hashUrl, originalUrl string) {
	query := fmt.Sprintf("INSERT INTO %s (short, original) VALUES ($1, $2)", tableName)
	hashUrl = port+hashUrl
	_ = r.db.QueryRow(query, hashUrl, originalUrl)
}

func (r *UrlRepos) GetOriginalUrl(shortUrl string) string {
	var originalUrl string
	query := fmt.Sprintf("SELECT original FROM %s WHERE short = $1", tableName)
	row, err := r.db.Query(query, shortUrl)
	if err != nil {
		return err.Error()
	}
	for row.Next() {
		err = row.Scan(&originalUrl)
		if err != nil {
			return err.Error()
		}
	}
	return originalUrl
}