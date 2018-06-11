package main

import (
	"encoding/json"
	"flag"
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
	return mux
}

func main() {
	port := *flag.String("port", "8000", "port on which to run the webserver")
	flag.Parse()
	log.Printf("webserver running on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, Init()))
}

func cool(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Response{true})
}
