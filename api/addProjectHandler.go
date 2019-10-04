package api

import (
  "net/http"

  "po_users/general"
  "po_users/database"
)

func addProjectHandler(w http.ResponseWriter, r *http.Request) {
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

  name, found := getQueryElement(query, "project_name")
  if !found {
    w.WriteHeader(400)
    return
  }

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
