package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var app App
var shortenedUrls map[string]ShortenedURL
var accessLogs map[string][]AccessLog
var dalMock DalMock

func TestMain(m *testing.M) {
	app = App{}
	shortenedUrls = make(map[string]ShortenedURL)
	accessLogs = make(map[string][]AccessLog)

	dalMock.ShortenedUrls = shortenedUrls
	dalMock.AccessLogs = accessLogs

	serviceParams := serviceParams{&dalMock, &dalMock, &dalMock, &dalMock}

	SeedData()

	app.Initialise(serviceParams)
	code := m.Run()
	defer os.Exit(code)
}

func SeedData() {
	shortenedUrls["poiuy"] = ShortenedURL{ShortCode: "poiuy", LongURL: "http://www.gdfgdfgfdfg.com/ertgfdgfdg"}
	shortenedUrls["gdfgd"] = ShortenedURL{ShortCode: "gdfgd", LongURL: "http://www.ljdkwdwed.com/1df1dfa6f"}
	shortenedUrls["jknba"] = ShortenedURL{ShortCode: "jknba", LongURL: "http://www.dgfdsvfdvf.com/qWQDED3CEC"}

	accessLogs["poiuy"] = append(accessLogs["poiuy"], AccessLog{DateTime: time.Now()})
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

func TestAccessingAShortenedURLIsLogged(t *testing.T) {
	// Arrange
	shortenedUrls["ABCDE"] = ShortenedURL{ShortCode: "ABCDE", LongURL: "http://www.google.co.uk/search?q=something+really+long"}

	req, _ := http.NewRequest("GET", "/r/ABCDE", nil)

	// Act
	response := executeRequest(req)

	// Assert
	checkResponseCode(t, http.StatusMovedPermanently, response.Code)

	val, ok := accessLogs["ABCDE"]
	if !ok {
		t.Error("Access was not logged")
	}

	for _, item := range val {
		fmt.Printf("%v", item.DateTime.String())
	}
}

func TestAccessURLThenReadAccessLogs(t *testing.T) {
	shortenedUrls["ABCDE"] = ShortenedURL{ShortCode: "ABCDE", LongURL: "http://www.google.co.uk/search?q=something+really+long"}
	urlReq, _ := http.NewRequest("GET", "/r/ABCDE", nil)
	_ = executeRequest(urlReq)

	now := time.Now()
	from := now.Add(time.Duration(-1) * time.Minute).Format(time.RFC3339)
	to := now.Add(time.Duration(1) * time.Minute).Format(time.RFC3339)

	accessReqURL := fmt.Sprintf("/api/access_logs/ABCDE?from=%v&to=%v", from, to)
	accessReq, _ := http.NewRequest("GET", accessReqURL, nil)
	accessResponse := executeRequest(accessReq)
	checkResponseCode(t, http.StatusOK, accessResponse.Code)
	fmt.Printf("\n%v\n", accessResponse.Body.String())
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
