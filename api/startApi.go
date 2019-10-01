package api

import (
  "net/http"

  "github.com/gorilla/mux"
)

func StartAPI(port string) {
  r := mux.NewRouter()
  r.HandleFunc("/create", createHandler).Methods("POST")

  r.HandleFunc("/find", findHandler).Methods("GET")
  http.Handle("/", r)

  http.ListenAndServe(port, nil)
}
