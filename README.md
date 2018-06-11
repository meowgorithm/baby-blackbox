[![GoDoc Badge](https://godoc.org/github.com/magicnumbers/blackbox?status.svg)](http://godoc.org/github.com/magicnumbers/blackbox)

Blackbox
========

Blackbox testing for Go JSON APIs. Currently, it supports the Go standard
library’s Multiplexer (`http.ServeMux`) and [Goji’s][goji] multiplexer
(`goji.Mux`) multiplexer.

[goji]: http://goji.io


## Pithy Example

```go

package main_test

import (
    "testing"
    "blackbox"
    app "."
)

type user struct {
    ID int `json:"id,omitempty"`
    Name string `json:"name"`
}

type apiError {
    Code int `json:"code"`
    Message string `json:"message"`
}

func TestMain(m *testing.M) {
    os.Exit(m.Run())
}

func TestStuff(t *testing.T) {

    // Create a blackbox testing thing from your application's multiplexer
    api := blackbox.New(t, app.GetMux())

    // Create a user, expecting to get an ID back
    u := user{Name: "Frankie"}
    api.Request("POST", "/user", u)
        .Created() // Assert that we want a 201 Created HTTP status
        .JSON(&u)  // Decode the response into a struct

    if u.ID == 0 {
        t.Error("expected to receive an ID, but we did not")
    }

    var apiErr apiError

    // Check another route, expecting an error
    api.Request("GET", "/cats", nil)
        .Status(http.StatusInternalServerError)
        .JSON(&apiErr)

    if err.Code != 21 {
        t.Errorf("thought we were gonna get error code 21, instead got %d", err.Code)
    }
}

```

For an example in the context of a more complete application, see the `example`
directory.


## License

MIT
