[![GoDoc Badge](https://godoc.org/github.com/magicnumbers/baby-blackbox?status.svg)](http://godoc.org/github.com/magicnumbers/baby-blackbox)

Baby Blackbox
=============

Blackbox testing for Go JSON APIs. Currently, it supports the Go standard
library’s Multiplexer (`http.ServeMux`) and the [Goji][goji] multiplexer
(`goji.Mux`).

[goji]: http://goji.io


## Example

```go

package main_test

import (
    "testing"
    "baby-blackbox"
    app "."
)

type user struct {
    ID int      `json:"id,omitempty"`
    Name string `json:"name"`
}

type apiError {
    Code int       `json:"code"`
    Message string `json:"message"`
}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}

func TestStuff(t *testing.T) {

    // Create a blackbox testing thing from your application's multiplexer
    api := blackbox.New(t, app.GetMux())

    // The payload we'll send in the next request
    u := user{Name: "Frankie"}

    // Create a user expecting to get an ID back in a JSON object. We assert
    // that we want a 201 Created status code.
    api.Request("POST", "/user", u).
        Created().
        JSON(&u)

    // Make sure we got an ID in the response
    if u.ID == 0 {
        t.Error("expected to receive an ID, but we did not")
    }

    var apiErr apiError

    // Check another route, expecting a 401 Unauthorized error
    api.Request("GET", "/cats", nil).
        Status(http.StatusUnauthorized).
        JSON(&apiErr)

    if err.Code != 21 {
        t.Errorf("thought we were gonna get error code 21, instead got %d", err.Code)
    }
}

```

For an example in the context of a more complete application, see the `example`
directory.


## License

MIT
