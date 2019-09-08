package main

import "time"

// ShortenURLRequest ...
type ShortenURLRequest struct {
	URLToShorten string `json:"urlToShorten"`
}

// ShortenURLResponse ...
type ShortenURLResponse struct {
	ShortenedURL string `json:"shortenedUrl"`
}

// AccessLogsResponse ...
type AccessLogsResponse struct {
	ShortCode string              `json:"shortCode"`
	Total     int                 `json:"total"`
	Logs      []AccessLogResponse `json:"logs"`
}

// AccessLogResponse ...
type AccessLogResponse struct {
	DateTime time.Time `json:"dateTime"`
}
