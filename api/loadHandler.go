package api

import (
  "net/http"
  "encoding/json"

  "po_users/general"
  "po_users/database"
)

type FindResponse struct {
  User general.ResponseUser `json:"user"`
}

func sendResult(resp FindResponse, w http.ResponseWriter) {
  jsonResponse, err := json.Marshal(resp)
  if err != nil {
    w.WriteHeader(400)

    return
  }

  w.WriteHeader(200)
  w.Write(jsonResponse)
}

func findWithSessionID(id string, w http.ResponseWriter) {
  res, err := database.FindUser_sessionID(id)
  if err != nil {
    w.WriteHeader(400)
    return
  }

  response := FindResponse{
    User: res.ConvertToResponse(),
  }

  sendResult(response, w)
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
  query := r.URL.Query()

  sessionID, sessionID_ok := query["sessionID"]
  if sessionID_ok && len(sessionID) > 0 {
    findWithSessionID(sessionID[0], w)
    return
  }


  w.WriteHeader(200)
}
