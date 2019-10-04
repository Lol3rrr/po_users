package api

import (
  "net/http"

  "po_users/database"
)

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
  query := getQuery(r)

  sessionID, found := getQueryElement(query, "sessionID")
  if !found {
    w.WriteHeader(400)
    return
  }

  id, found := getQueryElement(query, "project_id")
  if !found {
    w.WriteHeader(400)
    return
  }

  user, err := database.LoadUser_sessionID(sessionID)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  index := -1
  for i, val := range user.Projects {
    if val.ID == id {
      index = i
      break
    }
  }

  if index == -1 {
    w.WriteHeader(400)
    return
  }

  user.Projects = append(user.Projects[:index], user.Projects[index+1:]...)

  database.UpdateUser(user)

  w.WriteHeader(200)
}
