package main

import "time"

type shortenedURLAdder interface {
	AddShortenedURL(su ShortenedURL) error
}

type shortenedURLGetter interface {
	GetShortenedURL(shortCode string) (*ShortenedURL, error)
}

type accessLogger interface {
	LogAccess(shortCode string) error
}

type accessLogGetter interface {
	GetAccessLogs(shortCode string, from time.Time, to time.Time) ([]AccessLog, error)
}
