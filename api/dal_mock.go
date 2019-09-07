package main

import "github.com/stretchr/testify/mock"

type DalMock struct {
	mock.Mock
	ShortenedUrls map[string]ShortenedURL
}

func (m *DalMock) Add(su ShortenedURL) error {
	m.ShortenedUrls[su.ShortCode] = su
	return nil
}

func (m *DalMock) Get(shortCode string) (*ShortenedURL, error) {
	su := m.ShortenedUrls[shortCode]
	return &su, nil
}
