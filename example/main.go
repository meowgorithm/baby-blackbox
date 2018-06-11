package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

// Response is the only response the server returns. And it's a very, very cool
// response.
type Response struct {
	Cool bool `json:"cool"`
}

// Init performs initialization stuff. We export it so we can access it in
// blackbox testing.
func Init() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", cool)
	mux.HandleFunc("/notcool", notCool)
	return mux
}

// Our one and only handler
func cool(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{true})
}

func notCool(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprint(w, "You are not allowed to taste the forbidden fruit.")
}

func main() {
	port := *flag.String("port", "8000", "port on which to run the webserver")
	flag.Parse()
	log.Printf("webserver running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, Init()))
}
