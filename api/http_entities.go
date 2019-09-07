package main

// ShortenURLRequest ...
type ShortenURLRequest struct {
	URLToShorten string `json:"urlToShorten"`
}

// ShortenURLResponse ...
type ShortenURLResponse struct {
	ShortenedURL string `json:"shortenedUrl"`
}
