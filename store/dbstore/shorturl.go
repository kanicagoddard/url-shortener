package dbstore

import (
	"log/slog"
	"url-shortener/store"
)

type ShortURLStore struct {
	shortURLs []store.ShortURL // imports from store.gox
	logger    *slog.Logger
}

// dependancy injection
type NewShortURLStoreParams struct {
	Logger *slog.Logger
}

func NewShortURLStore(params NewShortURLStoreParams) *ShortURLStore {
	shortURLs := []store.ShortURL{}

	return &ShortURLStore{
		shortURLs: shortURLs,
		logger:    params.Logger,
	}
}

// create shortURL
func (s *ShortURLStore) CreateShortURL(params store.CreateShortURLParams) (store.ShortURL, error) {
	shortURL := store.ShortURL{
		Destination: params.Destination,
		Slug:        params.Slug,
		ID:          len(s.shortURLs), //returns a link to the slice using len()
	}

	s.shortURLs = append(s.shortURLs, shortURL)

	s.logger.Info("short URL created", slog.Any("values -> ", shortURL))
	return shortURL, nil

}

// get a short URL by slug
func (s *ShortURLStore) GetShortURLBySlug(slug string) (*store.ShortURL, error) {
	// use a for loop to search through the shortURLs, if found, return results or error

	for _, shortURL := range s.shortURLs {
		if shortURL.Slug == slug {
			result := shortURL
			return &result, nil
		}
	}

	return nil, store.ErrShortURLNotFound
}
