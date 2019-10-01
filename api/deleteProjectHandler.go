package api

import (
  "net/http"

  "po_users/database"
)

func deleteProjectHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  rawSessionID, sesssionID_ok := query["sessionID"]
  rawID, id_ok := query["project_id"]
  if !sesssionID_ok || len(rawSessionID) <= 0 {
    w.WriteHeader(400)
    return
  }
  if !id_ok || len(rawID) <= 0 {
    w.WriteHeader(400)
    return
  }

  sessionID := rawSessionID[0]
  id := rawID[0]

  user, err := database.FindUser_sessionID(sessionID)
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
