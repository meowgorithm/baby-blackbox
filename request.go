package blackbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"text/tabwriter"

	goji "goji.io"
)

// New instantiates a new API test object with a standard http.ServeMux
func New(t *testing.T, mux *http.ServeMux) APITest {
	return APITest{t: t, mux: mux}
}

// NewWithGoji instantiates a new API test object with a goji.Mux
func NewWithGoji(t *testing.T, mux *goji.Mux) APITest {
	return APITest{t: t, gojiMux: mux}
}

// APITest is a helper for running a series of API tests. Initialize it once
// with New (or NewWithGoji) then call `Request` to issue API requests
type APITest struct {
	t       *testing.T
	mux     *http.ServeMux
	gojiMux *goji.Mux
}

// Request makes a call to a REST API. It returns a Response struct which
// contains information about the response to the request, as well as methods
// for analyzing and working with the request data. See `request.go`.
func (a *APITest) Request(method string, path string, body interface{}) Response {
	var (
		b   []byte
		err error
	)

	// Marshal JSON body, if present
	if body != nil {
		if b, err = json.Marshal(body); err != nil {
			a.t.Errorf("error encoding JSON: %v", err)
		}
	}

	// Make the request
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(method, path, bytes.NewReader(b))

	// Which mux should we be using?
	if a.mux != nil {
		a.mux.ServeHTTP(recorder, request)
	} else if a.gojiMux != nil {
		a.gojiMux.ServeHTTP(recorder, request)
	} else {
		a.t.Error("no multiplexer set")
		return Response{}
	}

	// Response struct for evaluation later
	var response Response
	response.new(a.t, recorder)

	// Print out request/response details
	if os.Getenv("DEBUG") != "" {
		w := tabwriter.NewWriter(os.Stdout, 2, 0, 2, '.', 0)
		fmt.Fprintln(w, ".\t", "................")
		fmt.Fprintln(w, "Request: \t HTTP", method, path)
		if b != nil {
			fmt.Fprintln(w, "Request Body: \t", string(b))
		}
		fmt.Fprintln(w, "Response: \t", response.StatusCode, response.StatusText)
		for k, v := range recorder.HeaderMap {
			fmt.Fprintln(w, k+": \t", v)
		}
		if string(response.Body) != "" {
			fmt.Fprintln(w, "Response Body: \t", string(response.Body))
		}
		if err := w.Flush(); err != nil {
			a.t.Error("Could not flush tabwriter:", err.Error())
		}
	}

	return response
}
