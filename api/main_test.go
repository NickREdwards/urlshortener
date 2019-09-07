package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/NickREdwards/urlshortener/api/dal"
)

var app App
var shortenedUrls map[string]dal.ShortenedURL
var dalMock dal.DalMock

func TestMain(m *testing.M) {
	app = App{}
	shortenedUrls = make(map[string]dal.ShortenedURL)
	dalMock.ShortenedUrls = shortenedUrls

	app.Initialise(&dalMock, &dalMock)
	code := m.Run()
	defer os.Exit(code)
}

func TestCreateNewShortenedURL(t *testing.T) {
	var jsonStr = []byte(`{"urlToShorten":"http://www.google.co.uk/search?q=something+really+long"}`)

	req, _ := http.NewRequest("POST", "/api/create", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, "{}", response.Body.String())
}

func TestGetIndividualAccountEndpoint(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/accounts/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)
	checkResponseBody(t, "{\"id\":1,\"username\":\"Nick\"}\n", response.Body.String())
}

func TestGetIndividualAccountEndpointInvalidAccountReturns404(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/accounts/0", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
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
