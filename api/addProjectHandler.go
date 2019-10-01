package api

import (
  "net/http"

  "po_users/general"
  "po_users/database"
)

func addProjectHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  rawSessionID, sesssionID_ok := query["sessionID"]
  rawID, id_ok := query["project_id"]
  rawName, name_ok := query["project_name"]
  if !sesssionID_ok || len(rawSessionID) <= 0 {
    w.WriteHeader(400)
    return
  }
  if !id_ok || len(rawID) <= 0 {
    w.WriteHeader(400)
    return
  }
  if !name_ok || len(rawName) <= 0 {
    w.WriteHeader(400)
    return
  }

  sessionID := rawSessionID[0]
  id := rawID[0]
  name := rawName[0]

  user, err := database.FindUser_sessionID(sessionID)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  tmpProject := general.Project{
    ID: id,
    Name: name,
  }

  if user.Projects == nil {
    user.Projects = make([]general.Project, 0)
  }

  user.Projects = append(user.Projects, tmpProject)

  database.UpdateUser(user)

  w.WriteHeader(200)
}
