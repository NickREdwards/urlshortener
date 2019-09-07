package main

import (
	"testing"

	"github.com/NickREdwards/urlshortener/api/dal"

	"github.com/gorilla/mux"
)

func TestInitialise(t *testing.T) {
	// Arrange
	service := Service{}
	db := dal.DalMock{}

	// Act
	service.Initialise(new(mux.Router), &db, &db)

	// Assert
}
