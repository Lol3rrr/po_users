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
  query := r.URL.Query()

  rawName, name_ok := query["name"]
  if !name_ok || len(rawName) <= 0 {
    w.WriteHeader(400)
    return
  }

  name := rawName[0]
  id := guuid.New().String()
  sessionID := guuid.New().String()

  user := general.User{
    ID: id,
    SessionID: sessionID,
    Name: name,
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
