package main

import (
	"time"

	"github.com/stretchr/testify/mock"
)

// DalMock allows the mocking out of the data access layer
type DalMock struct {
	mock.Mock
	ShortenedUrls map[string]ShortenedURL
	AccessLogs    map[string][]AccessLog
}

// AddShortenedURL to the mocked out list of shortened URLs
func (m *DalMock) AddShortenedURL(su ShortenedURL) error {
	m.ShortenedUrls[su.ShortCode] = su
	return nil
}

// GetShortenedURL returns a previously shortened URL by the short code
func (m *DalMock) GetShortenedURL(shortCode string) (*ShortenedURL, error) {
	su := m.ShortenedUrls[shortCode]
	return &su, nil
}

// LogAccess records access to a shortened URL
func (m *DalMock) LogAccess(shortCode string) error {
	if val, ok := m.AccessLogs[shortCode]; ok {
		val = append(val, AccessLog{time.Now().UTC()})
		return nil
	}

	m.AccessLogs[shortCode] = make([]AccessLog, 1)
	m.AccessLogs[shortCode] = append(m.AccessLogs[shortCode], AccessLog{time.Now().UTC()})

	return nil
}

// GetAccessLogs returns a slice of AccessLogs for the given shortCode, and in the given time frame
func (m *DalMock) GetAccessLogs(shortCode string, from time.Time, to time.Time) ([]AccessLog, error) {
	return nil, nil
}
