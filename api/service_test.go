package main

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestInitialise(t *testing.T) {
	// Arrange
	service := Service{}
	db := DalMock{}

	// Act
	service.Initialise(new(mux.Router), &db, &db)

	// Assert
}
