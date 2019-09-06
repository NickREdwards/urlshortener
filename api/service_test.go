package main

import (
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
)

var shortenedUrls map[string]ShortenedURL
var service *Service
var adder *shortURLAdderMock
var getter *shortURLGetterMock

func Setup() {
	service = new(Service)
	adder = new(shortURLAdderMock)
	getter = new(shortURLGetterMock)
	shortenedUrls = make(map[string]ShortenedURL)
}

type shortURLAdderMock struct {
	mock.Mock
}

func (m *shortURLAdderMock) Add(su ShortenedURL) error {
	shortenedUrls[su.shortCode] = su
	return nil
}

type shortURLGetterMock struct {
	mock.Mock
}

func (m *shortURLGetterMock) Get(shortCode string) (ShortenedURL, error) {
	return shortenedUrls[shortCode], nil
}

func TestInitialise(t *testing.T) {
	// Arrange
	Setup()

	// Act
	service.Initialise(new(mux.Router), adder, getter)

	// Assert
}
