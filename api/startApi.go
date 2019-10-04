package api

import (
  "net/http"

  "github.com/gorilla/mux"
)

func StartAPI(port string) {
  r := mux.NewRouter()
  r.HandleFunc("/create", createHandler).Methods("POST")

  r.HandleFunc("/projects/add", addProjectHandler).Methods("POST")
  r.HandleFunc("/projects/delete", deleteProjectHandler).Methods("POST")

  r.HandleFunc("/edit/name", editNameHandler).Methods("POST")

  r.HandleFunc("/load", loadHandler).Methods("GET")
  http.Handle("/", r)

  http.ListenAndServe(port, nil)
}
