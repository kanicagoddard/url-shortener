package dbstore

import (
	"log/slog"
	"url-shortener/store"
)

type ShortURLSTore struct {
	shortURLs []store.ShortURL
	logger    *slog.Logger
}

// dependancy injection
type NewShortURLStoreParams struct {
	Logger *slog.Logger
}

func NewShortURLStore(params NewShortURLStoreParams) *ShortURLSTore {
	shortURLs := []store.ShortURL{}

	return &ShortURLSTore{
		shortURLs: shortURLs,
		logger:    params.Logger,
	}
}
