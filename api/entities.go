package main

// ShortenedURL ...
type ShortenedURL struct {
	shortCode string
	longURL   string
}

// ShortenURLRequest ...
type ShortenURLRequest struct {
	URLToShorten string `json:"urlToShorten"`
}

// ShortenURLResponse ...
type ShortenURLResponse struct {
	ShortenedURL string `json:"shortenedUrl"`
}
