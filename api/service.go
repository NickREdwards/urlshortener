package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/NickREdwards/urlshortener/api/dal"
	"github.com/gorilla/mux"
)

const (
	shortCodeLength = 6
)

// Service - manages all routes and calls to the accounts resource
type Service struct {
	urlAdder  dal.ShortenedURLAdder
	urlGetter dal.ShortenedURLGetter
}

type apiError struct {
	Error string `json:"error"`
}

// Initialise - initialises list and creates default data
func (s *Service) Initialise(router *mux.Router, urlAdder dal.ShortenedURLAdder, urlGetter dal.ShortenedURLGetter) {
	s.registerRoutes(router)
	s.urlAdder = urlAdder
	s.urlGetter = urlGetter
}

func (s *Service) registerRoutes(router *mux.Router) {
	router.HandleFunc("/api/create", s.createShortenedURL).Methods("POST")

	router.HandleFunc(fmt.Sprintf("/r/{shortCode:(?:[A-Za-z0-9]{%v})}", shortCodeLength), s.resolveShortenedURL).Methods("GET")
	router.HandleFunc("/health", func(w http.ResponseWriter, req *http.Request) { io.WriteString(w, "healthy") }).Methods("GET")
}

func (s *Service) createShortenedURL(w http.ResponseWriter, r *http.Request) {
	var urlRequest ShortenURLRequest
	_ = json.NewDecoder(r.Body).Decode(&urlRequest)
	shortCode := NewShortCode(shortCodeLength)
	shortenedURL := dal.ShortenedURL{ShortCode: shortCode, LongURL: urlRequest.URLToShorten}
	err := s.urlAdder.Add(shortenedURL)
	if err != nil {
		returnError(w, errors.New("Error creating new shortened URL"))
	} else {
		returnJSON(w, shortenedURL)
	}
}

func (s *Service) resolveShortenedURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]
	su, err := s.urlGetter.Get(shortCode)
	if err != nil {
		returnError(w, errors.New("Error resolving shortend URL"))
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
