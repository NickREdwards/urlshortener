package main

// import (
// 	"net/http"
// 	"net/http/httptest"
// 	"os"
// 	"testing"
// )

// var app App

// func TestMain(m *testing.M) {
// 	app = App{}
// 	app.Initialise()
// 	code := m.Run()
// 	defer os.Exit(code)
// }

// func TestGetAllAccountsEndpoint(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/api/accounts", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)
// 	checkResponseBody(t, "[{\"id\":1,\"username\":\"Nick\"},{\"id\":2,\"username\":\"Jack\"},{\"id\":3,\"username\":\"Jake\"},{\"id\":4,\"username\":\"Matt\"}]\n", response.Body.String())
// }

// func BenchmarkGetAllAccountsEndpoint(t *testing.B) {
// 	req, _ := http.NewRequest("GET", "/api/accounts", nil)
// 	executeRequest(req)
// }

// func TestGetIndividualAccountEndpoint(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/api/accounts/1", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusOK, response.Code)
// 	checkResponseBody(t, "{\"id\":1,\"username\":\"Nick\"}\n", response.Body.String())
// }

// func TestGetIndividualAccountEndpointInvalidAccountReturns404(t *testing.T) {
// 	req, _ := http.NewRequest("GET", "/api/accounts/0", nil)
// 	response := executeRequest(req)
// 	checkResponseCode(t, http.StatusNotFound, response.Code)
// }

// func executeRequest(req *http.Request) *httptest.ResponseRecorder {
// 	rr := httptest.NewRecorder()
// 	app.router.ServeHTTP(rr, req)

// 	return rr
// }
// func checkResponseCode(t *testing.T, expected, actual int) {
// 	if expected != actual {
// 		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
// 	}
// }

// func checkResponseBody(t *testing.T, expected, actual string) {
// 	if expected != actual {
// 		t.Errorf("Expected body %q. Got %q\n", expected, actual)
// 	}
// }
