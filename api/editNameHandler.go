package api

import (
  "net/http"

  "po_users/database"
)

func editNameHandler(w http.ResponseWriter, r *http.Request) {
  query := getQuery(r)

  sessionID, found := getQueryElement(query, "sessionID")
  if !found {
    w.WriteHeader(400)
    return
  }

  nName, found := getQueryElement(query, "name")
  if !found || nName == "" {
    w.WriteHeader(400)
    return
  }

  user, err := database.LoadUser_sessionID(sessionID)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  user.Name = nName

  database.UpdateUser(user)

  w.WriteHeader(200)
}
