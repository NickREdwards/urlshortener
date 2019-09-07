package dal

import "github.com/stretchr/testify/mock"

type DalMock struct {
	mock.Mock
	ShortenedUrls map[string]ShortenedURL
}

func (m *DalMock) Add(su ShortenedURL) error {
	m.ShortenedUrls[su.ShortCode] = su
	return nil
}

func (m *DalMock) Get(shortCode string) (ShortenedURL, error) {
	return m.ShortenedUrls[shortCode], nil
}
