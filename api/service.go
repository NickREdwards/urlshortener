package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	shortCodeLength = 6
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
	router.HandleFunc("/r/{shortCode}", s.resolveShortenedURL).Methods("GET")
}

func (s *Service) createShortenedURL(w http.ResponseWriter, r *http.Request) {
	var urlRequest ShortenURLRequest
	_ = json.NewDecoder(r.Body).Decode(&urlRequest)
	shortCode := NewShortCode(shortCodeLength)
	shortenedURL := ShortenedURL{ShortCode: shortCode, LongURL: urlRequest.URLToShorten}
	err := s.urlAdder.Add(shortenedURL)
	if err != nil {
		returnError(w, errors.New("Error creating new shortened URL"))
	} else {
		url := fmt.Sprintf("%v/r/%v", r.Host, shortenedURL.ShortCode)
		response := ShortenURLResponse{ShortenedURL: url}
		returnJSON(w, response)
	}
}

func (s *Service) resolveShortenedURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]
	su, err := s.urlGetter.Get(shortCode)
	if err != nil {
		returnError(w, errors.New("Error resolving shortened URL"))
	} else {
		http.Redirect(w, r, su.LongURL, 301)
	}
}

func returnJSON(w http.ResponseWriter, o interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(o)
}

func returnError(w http.ResponseWriter, err error) {
	returnJSON(w, apiError{fmt.Sprintf("%v", err)})
}
