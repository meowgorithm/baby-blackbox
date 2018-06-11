package blackbox

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// Response contains helper members and methods for working with API tests
type Response struct {
	t          *testing.T
	StatusCode int
	StatusText string
	Body       []byte
}

// Initialize a new response object. This is called interally with
// `APITest.Request()`
func (a *Response) new(t *testing.T, r *httptest.ResponseRecorder) {
	var err error

	a.t = t
	a.StatusCode = r.Code
	a.StatusText = http.StatusText(r.Code)

	// Read response body
	if a.Body, err = ioutil.ReadAll(r.Body); err != nil {
		t.Errorf("could not read response body: %s", err.Error())
	}

	// Clean up body
	a.Body = []byte(strings.Trim(string(a.Body), "\n "))
}

// Cool checks weather HTTP status code is in the 200s-300s range (i.e. not an
// error)
func (a Response) Cool() Response {
	// In other words, if status is not in the 200s-300s range
	if a.StatusCode < http.StatusOK || a.StatusCode >= http.StatusBadRequest {
		a.t.Errorf(fmt.Sprintf("HTTP status not ok: %d %s", a.StatusCode, a.StatusText))
	}
	return a
}

// Status checks that the HTTP status code matches an expected one
func (a Response) Status(code int) Response {
	if a.StatusCode != code {
		a.t.Errorf("HTTP: expected %d %s, got %d %s", code,
			http.StatusText(code), a.StatusCode, a.StatusText)
	}
	return a
}

// OK is a shortcut for checking for an HTTP 200 response
func (a Response) OK() Response {
	a.Status(http.StatusOK)
	return a
}

// Created is a shortcut for checking for an HTTP 201 response
func (a Response) Created() Response {
	a.Status(http.StatusCreated)
	return a
}

// NoContent is a shortcut for checking for an HTTP 204 response
func (a Response) NoContent() Response {
	a.Status(http.StatusNoContent)
	return a
}

// NotFound is a shortcut for checking for an HTTP 404 response
func (a Response) NotFound() Response {
	a.Status(http.StatusNotFound)
	return a
}

// InternalServerError is a shortcut for checking for an HTTP 500 response
func (a Response) InternalServerError() Response {
	a.Status(http.StatusInternalServerError)
	return a
}

// JSON decodes the request body into the given interface
func (a Response) JSON(i interface{}) Response {
	if err := json.Unmarshal(a.Body, &i); err != nil {
		a.t.Errorf(fmt.Sprintf("error decoding JSON: %v", err))
	}
	return a
}
