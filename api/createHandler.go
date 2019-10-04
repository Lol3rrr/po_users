package api

import (
  "net/http"
  "encoding/json"

  guuid "github.com/google/uuid"

  "po_users/general"
  "po_users/database"
)

type CreateResponse struct {
  SessionID string `json:"sessionID"`
}

func createHandler(w http.ResponseWriter, r *http.Request) {
  query := getQuery(r)

  name, found := getQueryElement(query, "name")
  if !found {
    w.WriteHeader(400)
    return
  }
  googleId, _ := getQueryElement(query, "id_google")

  id := guuid.New().String()
  sessionID := guuid.New().String()

  user := general.User{
    ID: id,
    SessionID: sessionID,
    Name: name,
    GoogleID: googleId,
  }

  exists, existingUser := database.UserExists(user)
  if exists {
    user = existingUser
    user.SessionID = sessionID
  }

  database.UpdateUser(user)

  resp := CreateResponse{
    SessionID: sessionID,
  }

  jsonResponse, err := json.Marshal(resp)
  if err != nil {
    w.WriteHeader(400)

    return
  }

  w.WriteHeader(200)
  w.Write(jsonResponse)
}
