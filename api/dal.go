package main

// ShortenedURLAdder ...
type ShortenedURLAdder interface {
	Add(su ShortenedURL) error
}

// ShortenedURLGetter ...
type ShortenedURLGetter interface {
	Get(shortCode string) (*ShortenedURL, error)
}
