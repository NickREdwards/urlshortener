package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type serviceParams struct {
	urlAdder        shortenedURLAdder
	urlGetter       shortenedURLGetter
	accessLogger    accessLogger
	accessLogGetter accessLogGetter
}

// Service ...
type Service struct {
	urlAdder        shortenedURLAdder
	urlGetter       shortenedURLGetter
	accessLogger    accessLogger
	accessLogGetter accessLogGetter
}

type apiError struct {
	Error string `json:"error"`
}

// Initialise ...
func (s *Service) Initialise(router *mux.Router, params serviceParams) {
	s.registerRoutes(router)
	s.urlAdder = params.urlAdder
	s.urlGetter = params.urlGetter
	s.accessLogger = params.accessLogger
	s.accessLogGetter = params.accessLogGetter
}

func (s *Service) registerRoutes(router *mux.Router) {
	router.HandleFunc("/api/create", s.createShortenedURL).Methods("POST")
	router.HandleFunc("/api/access_logs/{shortCode}", s.getAccessLogs).Methods("GET")

	router.HandleFunc("/r/{shortCode}", s.resolveShortenedURL).Methods("GET")
}

func (s *Service) createShortenedURL(w http.ResponseWriter, r *http.Request) {
	var urlRequest ShortenURLRequest
	_ = json.NewDecoder(r.Body).Decode(&urlRequest)
	shortCode := NewShortCode(shortCodeLength)
	shortenedURL := ShortenedURL{ShortCode: shortCode, LongURL: urlRequest.URLToShorten}
	err := s.urlAdder.AddShortenedURL(shortenedURL)
	if err != nil {
		returnError(w, errors.New("Error creating new shortened URL"))
	} else {
		url := fmt.Sprintf("%v/r/%v", r.Host, shortenedURL.ShortCode)
		response := ShortenURLResponse{ShortenedURL: url}
		returnJSON(w, response)
	}
}

func (s *Service) getAccessLogs(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	shortCode := vars["shortCode"]
	var timeFrom time.Time
	var timeTo = time.Now()
	var err error

	query := r.URL.Query()
	from := query.Get("from")
	to := query.Get("to")

	if from != "" {
		timeFrom, err = time.Parse("2006-01-02 15:04:05", from)
		if err != nil {
			returnError(w, errors.New("Invalid value supplied for 'from'"))
			return
		}
	}

	if to != "" {
		timeTo, err = time.Parse("2006-01-02 15:04:05", to)
		if err != nil {
			returnError(w, errors.New("Invalid value supplied for 'to'"))
			return
		}
	}

	logs, err := s.accessLogGetter.GetAccessLogs(shortCode, timeFrom, timeTo)
	if err != nil {
		returnError(w, err)
	} else {
		response := AccessLogsResponse{}
		response.ShortCode = shortCode
		response.Total = len(logs)
		response.Logs = make([]AccessLogResponse, response.Total)
		for i, l := range logs {
			response.Logs[i] = AccessLogResponse{l.DateTime}
		}
		returnJSON(w, response)
	}
}

func (s *Service) resolveShortenedURL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	su, err := s.urlGetter.GetShortenedURL(shortCode)

	if err != nil {
		returnError(w, errors.New("Error resolving shortened URL"))
	} else {
		s.accessLogger.LogAccess(shortCode)

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
