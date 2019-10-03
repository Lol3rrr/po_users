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

func getGoogleID(query map[string][]string) (string) {
  rawID, ok := query["id_google"]
  if !ok || len(rawID) <= 0 {
    return ""
  }

  return rawID[0]
}

func createHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  rawName, name_ok := query["name"]
  if !name_ok || len(rawName) <= 0 {
    w.WriteHeader(400)
    return
  }

  googleId := getGoogleID(query)
  name := rawName[0]
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
