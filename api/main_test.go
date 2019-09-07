package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var app App
var shortenedUrls map[string]ShortenedURL
var dalMock DalMock

func TestMain(m *testing.M) {
	app = App{}
	shortenedUrls = make(map[string]ShortenedURL)
	dalMock.ShortenedUrls = shortenedUrls

	app.Initialise(&dalMock, &dalMock)
	code := m.Run()
	defer os.Exit(code)
}

func TestCreateNewShortenedURLRunsSuccessfully(t *testing.T) {
	// Arrange
	var jsonStr = []byte(`{"urlToShorten":"http://www.google.co.uk/search?q=something+really+long"}`)

	req, _ := http.NewRequest("POST", "/api/create", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	// Act
	response := executeRequest(req)

	// Assert
	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestResolveShortenedURL(t *testing.T) {
	// Arrange
	shortenedUrls["ABCDE"] = ShortenedURL{ShortCode: "ABCDE", LongURL: "http://www.google.co.uk/search?q=something+really+long"}

	req, _ := http.NewRequest("GET", "/r/ABCDE", nil)

	// Act
	response := executeRequest(req)

	// Assert
	checkResponseCode(t, http.StatusMovedPermanently, response.Code)
}

func TestCreateShortenedURLThenResolveIt(t *testing.T) {
	var jsonStr = []byte(`{"urlToShorten":"http://www.google.co.uk/search?q=something+really+really+really+long"}`)

	createReq, _ := http.NewRequest("POST", "/api/create", bytes.NewBuffer(jsonStr))
	createReq.Header.Set("Content-Type", "application/json")
	createResponse := executeRequest(createReq)
	checkResponseCode(t, http.StatusOK, createResponse.Code)

	var shortenedURL ShortenURLResponse
	json.Unmarshal([]byte(createResponse.Body.String()), &shortenedURL)

	assert.NotNil(t, shortenedURL)

	getReq, _ := http.NewRequest("GET", shortenedURL.ShortenedURL, nil)
	getResponse := executeRequest(getReq)
	checkResponseCode(t, http.StatusMovedPermanently, getResponse.Code)
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.router.ServeHTTP(rr, req)

	return rr
}
func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

func checkResponseBody(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("Expected body %q. Got %q\n", expected, actual)
	}
}
