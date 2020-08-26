package main

import (
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

func Route(r *mux.Router) {
    r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "Hello world!")
    })
}
