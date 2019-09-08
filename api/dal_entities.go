package main

import "time"

// ShortenedURL ...
type ShortenedURL struct {
	ShortCode string
	LongURL   string
}

// AccessLog ...
type AccessLog struct {
	DateTime time.Time
}
