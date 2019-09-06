package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

// Service - manages all routes and calls to the accounts resource
type Service struct {
	urlAdder  ShortenedURLAdder
	urlGetter ShortenedURLGetter
}

type apiError struct {
	Error string `json:"error"`
}

// Initialise - initialises list and creates default data
func (s *Service) Initialise(router *mux.Router, urlAdder ShortenedURLAdder, urlGetter ShortenedURLGetter) {
	s.registerRoutes(router)
	s.urlAdder = urlAdder
	s.urlGetter = urlGetter
}

func (s *Service) registerRoutes(router *mux.Router) {
	router.HandleFunc("/api/create", s.createShortenedURL).Methods("POST")
	router.HandleFunc("/api/health", func(w http.ResponseWriter, req *http.Request) { io.WriteString(w, "healthy") })
}

func (s *Service) createShortenedURL(w http.ResponseWriter, r *http.Request) {
	var urlRequest ShortenURLRequest
	_ = json.NewDecoder(r.Body).Decode(&urlRequest)
	shortCode := NewShortCode()
	shortenedURL := ShortenedURL{shortCode, urlRequest.URLToShorten}
	err := s.urlAdder.Add(shortenedURL)
	if err != nil {
		returnError(w, errors.New("Error creating new shortened URL"))
	} else {
		returnJSON(w, shortenedURL)
	}
}

func returnJSON(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(o)
}

func returnError(w http.ResponseWriter, err error) {
	returnJSON(w, apiError{fmt.Sprintf("%v", err)})
}
