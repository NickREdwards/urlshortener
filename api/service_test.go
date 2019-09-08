package main

import (
	"testing"

	"github.com/gorilla/mux"
)

func TestInitialise(t *testing.T) {
	// Arrange
	service := Service{}
	db := DalMock{}
	serviceParams := serviceParams{&db, &db, &db, &db}

	// Act
	service.Initialise(new(mux.Router), serviceParams)

	// Assert
}
