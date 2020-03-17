package goolpie

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var stubConfig = Config{
	BasePath: "/v1",
	Endpoints: []EndpointSettings{
		{
			Endpoint:     "/test",
			Method:       "GET",
			ResponseCode: http.StatusOK,
			ResponseBody: "Success",
		},
		{
			Endpoint:     "/other",
			Method:       "POST",
			Request:      "test",
			ResponseCode: http.StatusOK,
			ResponseBody: "Success",
		},
		{
			Endpoint:     "/fail",
			Method:       "PUT",
			Request:      "test",
			ResponseCode: http.StatusMethodNotAllowed,
			ResponseBody: "Fail",
		},
	},
}

func TestStubServer(t *testing.T) {
	testCases := []struct {
		endpoint     string
		method       string
		request      string
		response     string
		responseCode int
	}{
		{
			endpoint:     "/v1/test",
			method:       http.MethodGet,
			request:      "",
			response:     "Success",
			responseCode: http.StatusOK,
		},
		{
			endpoint:     "/v1/other",
			method:       http.MethodPost,
			request:      "test",
			response:     "Success",
			responseCode: http.StatusOK,
		},
		{
			endpoint:     "/v1/fail",
			method:       http.MethodPut,
			request:      "test",
			response:     "Fail",
			responseCode: http.StatusMethodNotAllowed,
		},
	}

	stubServer := NewStubServer(stubConfig)

	for _, tt := range testCases {
		request, _ := http.NewRequest(tt.method, tt.endpoint, strings.NewReader(tt.request))
		response := httptest.NewRecorder()

		stubServer.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), tt.response)
		assertResponseCode(t, response.Code, tt.responseCode)
	}
}

func assertResponseBody(t *testing.T, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("Response is not correct. Expected: %s, Got: %s", expected, got)
	}
}

func assertResponseCode(t *testing.T, got, expected int) {
	t.Helper()
	if got != expected {
		t.Errorf("Response status is not correct. Expected: %d, Got: %d", expected, got)
	}
}
